package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetQuestionRaw(token string, questionId int) (string, error) {
	url := fmt.Sprintf("https://binarysearch.com/api/publicquestions/?questionIds=%d", questionId)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("authority", "binarysearch.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cookie", "_ga=GA1.2.367669984.1659692871; _gid=GA1.2.1606135968.1661132149; _gat=1")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://binarysearch.com/problems/Flipped-Matrix")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"104\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"104\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	req.Header.Add("x-access-token", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}

func QueryQuestionList(token string, term *string) (string, error) {

	url := "https://binarysearch.com/api/questionlist"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	q := req.URL.Query()
	q.Add("list", "")
	q.Add("orderBy", "history")
	q.Add("page", "0")
	q.Add("term", *term)
	req.URL.RawQuery = q.Encode()

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("authority", "binarysearch.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cookie", "_ga=GA1.2.367669984.1659692871; _gid=GA1.2.1606135968.1661132149; _gat=1")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"104\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"104\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	req.Header.Add("x-access-token", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}
