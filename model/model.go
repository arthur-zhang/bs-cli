package model

type Question struct {
	Id                  int         `json:"id"`
	Title               string      `json:"title"`
	Status              string      `json:"status"`
	Contest             bool        `json:"contest"`
	Content             string      `json:"content"`
	Clarification       interface{} `json:"clarification"`
	SolutionExplanation string      `json:"solutionExplanation"`
	Background          interface{} `json:"background"`
	Difficulty          int         `json:"difficulty"`
	Attempted           int         `json:"attempted"`
	Solved              int         `json:"solved"`
	Level               int         `json:"level"`
	TimeLimit           int         `json:"timeLimit"`
	InputType           struct {
		Type  string `json:"type"`
		Items []struct {
			Type  string `json:"type"`
			Items struct {
				Type string `json:"type"`
			} `json:"items"`
			Title string `json:"title"`
		} `json:"items"`
	} `json:"inputType"`
	OutputType struct {
		Type string `json:"type"`
	} `json:"outputType"`
	Format     interface{} `json:"format"`
	Slug       string      `json:"slug"`
	CreateTime int         `json:"createTime"`
	UpdateTime int         `json:"updateTime"`

	Topics    []interface{} `json:"topics"`
	Companies []interface{} `json:"companies"`
	Hints     []struct {
		Id         int    `json:"id"`
		QuestionId int    `json:"questionId"`
		Content    string `json:"content"`
		Order      int    `json:"order"`
		UserId     int    `json:"userId"`
		Public     bool   `json:"public"`
	} `json:"hints"`
	Testcases []struct {
		Id             int           `json:"id"`
		Input          []interface{} `json:"input"`
		ExpectedOutput interface{}   `json:"expectedOutput"`
		QuestionId     int           `json:"questionId"`
		IsExample      bool          `json:"isExample"`
		Explanation    string        `json:"explanation"`
		Methods        interface{}   `json:"methods"`
		Special        bool          `json:"special"`
	} `json:"testcases"`
	InputParameters []string `json:"inputParameters"`
	Boilerplate     struct {
		Python     string `json:"python"`
		Ruby       string `json:"ruby"`
		Javascript string `json:"javascript"`
		Typescript string `json:"typescript"`
		Php        string `json:"php"`
		Java       string `json:"java"`
		Cpp        string `json:"cpp"`
	} `json:"boilerplate"`
	Form struct {
		Title   string `json:"title"`
		Methods []struct {
			Title     string `json:"title"`
			InputType []struct {
				Type  string `json:"type"`
				Items struct {
					Type string `json:"type"`
				} `json:"items"`
				Title string `json:"title"`
			} `json:"inputType"`
			OutputType struct {
				Type string `json:"type"`
			} `json:"outputType"`
		} `json:"methods"`
	} `json:"form"`
	IsInteractive bool `json:"isInteractive"`
}
type QuestionList struct {
	Questions []struct {
		Id         int    `json:"id"`
		Slug       string `json:"slug"`
		Title      string `json:"title"`
		Difficulty int    `json:"difficulty"`
	} `json:"questions"`
	Pages int `json:"pages"`
}
