package main

import (
	"bs_cli/api"
	"bs_cli/model"
	"bytes"

	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	AccessToken string
	DevDir      string
}

var difficulty = []string{
	"Easy",
	"Medium",
	"Hard",
	"Harder",
}

func beautify(question *model.Question) string {

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("[%d] %s [%s]\n", question.Id, question.Title, difficulty[question.Difficulty]))
	buffer.WriteString("\n")
	buffer.WriteString(question.Content)
	buffer.WriteString("\n")
	for i := range question.Testcases {
		buffer.WriteString("\n--------------------------------------------------\n")
		buffer.WriteString(fmt.Sprintf("[TestCase %d]\n", i))
		testCaseItem := question.Testcases[i]
		input := testCaseItem.Input
		buffer.WriteString(fmt.Sprintf("Input\n"))
		for inputIdx := range question.InputParameters {
			str, err := json.Marshal(input[inputIdx])
			if err != nil {
				panic(err)
			}
			buffer.WriteString(fmt.Sprintf("\t%s = %s\n", question.InputParameters[inputIdx], str))
		}
		buffer.WriteString("\n")

		output := testCaseItem.ExpectedOutput
		outputStr, err := json.Marshal(output)
		if err != nil {
			return ""
		}
		buffer.WriteString(fmt.Sprintf("Expect Output: %s\n", string(outputStr)))
		buffer.WriteString(fmt.Sprintf("Explanation: %s\n", testCaseItem.Explanation))
		buffer.WriteString("--------------------------------------------------\n")
	}
	return buffer.String()
}
func processQuestionList(conf *Config, arg *string) {
	questionName := strings.ReplaceAll(*arg, "-", " ")
	jsonStr, err := api.QueryQuestionList(conf.AccessToken, &questionName)
	questionList := model.QuestionList{}
	err = json.Unmarshal([]byte(jsonStr), &questionList)
	if err != nil {
		return
	}
	for i := range questionList.Questions {
		questionItem := questionList.Questions[i]
		fmt.Printf("[%d] %s\n", questionItem.Id, questionItem.Title)
	}

}
func processQuestion(conf *Config, arg int) {
	jsonStr, err := api.GetQuestionRaw(conf.AccessToken, arg)
	if err != nil {
		panic(err.Error())
	}
	_ = ioutil.WriteFile("/tmp/output.json", []byte(jsonStr), 0644)

	var questionArr []model.Question
	err = json.Unmarshal([]byte(jsonStr), &questionArr)
	if err != nil {
		return
	}
	question := questionArr[0]
	questionDesc := beautify(&question)
	println(questionDesc)

	parts := strings.Split(question.Title, " ")
	dirName := strings.Join(parts, "-")

	dir := fmt.Sprintf("%s/%d.%s", conf.DevDir, question.Id, dirName)

	genCpp(dir, &question)
	genDescTxt(dir, questionDesc, &question)
}
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
func main() {
	noFlag := flag.Int("s", 0, "Input Problem number")
	term := flag.String("l", "Add", "Input Problem name")
	flag.Parse()
	if isFlagPassed("s") {
	} else if isFlagPassed("l") {
	} else {
		fmt.Println("Usage: bs [-ls] params...")
		flag.PrintDefaults()
		os.Exit(1)
	}
	var conf Config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	configDir := fmt.Sprintf("%s/.bs", homeDir)

	_, err = os.Stat(configDir)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		err = os.Mkdir(configDir, 0755)
		if err != nil {
			panic("config dir mkdir failed")
			return
		}
	}

	tomlDataBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/conf.toml", configDir))
	if err != nil {
		panic("${HOME}/.bs/conf.toml read failed")
		return
	}
	tomlData := string(tomlDataBytes)
	_, err = toml.Decode(tomlData, &conf)
	if err != nil {
		return
	}
	if isFlagPassed("s") {
		processQuestion(&conf, *noFlag)
		return
	}
	if isFlagPassed("l") {
		processQuestionList(&conf, term)
		return
	}

}

func genDescTxt(dirPath string, content string, question *model.Question) {

	descFilePath := fmt.Sprintf("%s/%d.txt", dirPath, question.Id)

	file, err := os.OpenFile(descFilePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return
	}
}
func genCpp(dir string, question *model.Question) {
	//dir := fmt.Sprintf("%s/%d.%s", conf.DevDir, question.Id, dirName)
	err := os.Mkdir(dir, 0755)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	parts := strings.Split(question.Title, " ")
	dirName := strings.Join(parts, "-")
	cppFileName := fmt.Sprintf("%d_%s.cpp", question.Id, dirName)
	cppFilePath := fmt.Sprintf("%s/%s", dir, cppFileName)
	file, err := os.OpenFile(cppFilePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("#include <gtest/gtest.h>\n#include <iostream>\n#include <vector>\n#include <cmath>\n#include <queue>\n#include <unordered_set>\n#include <unordered_map>\n\nusing namespace std;\n")
	writer.WriteString(fmt.Sprintf("namespace bs_%d {\n", question.Id))
	writer.WriteString("\n")
	writer.WriteString(question.Boilerplate.Cpp)
	writer.WriteString("\n")
	writer.WriteString(fmt.Sprintf("\tTEST(bs_%d, bs_solotion_%d) {\n", question.Id, question.Id))
	writer.WriteString("        cout << \"enter test\" << endl;\n    }\n}")
	writer.Flush()

}
