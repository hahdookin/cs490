package util

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
