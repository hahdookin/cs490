package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// http://ec2-3-92-132-35.compute-1.amazonaws.com/studentexamanswers
// GET: http://ec2-3-92-132-35.compute-1.amazonaws.com/questions

// Questions (Database)
type DBQuestion_Test struct {
	Arguments []string
	Output    string
}

type DBQuestion struct {
	ID           int
	Title        string
	Desc         string
	Difficulty   string
	FunctionName string
	Parameters   []string
	Tests        []DBQuestion_Test
}

// StudentExamAnswers
type StudentExamAnswer_Test struct {
	Pass          bool
	Points        int
	StudentOutput string
}

type StudentExamAnswer struct {
	QID                 int
	StudentExamAnswerID int
	Code                string
	Tests               []StudentExamAnswer_Test
	Comment             string
	Id                  int
}

// DBGetJSON: Get's Info from DB and converts it into a DBQuestion Struct
func DBGetJSON(endpoint string) DBQuestion {
	var backendQuest DBQuestion
	resp, err := http.Get(endpoint)
	Check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)

	err = json.Unmarshal(body, &backendQuest)
	Check(err)

	return backendQuest
}
