package util

// http://ec2-3-92-132-35.compute-1.amazonaws.com/studentexamanswers
// GET: http://ec2-3-92-132-35.compute-1.amazonaws.com/questions

type DBQuestion_Test struct {
	Arguments []string
	Output    string
}

type DBQuestions struct {
	ID           int
	Title        string
	Desc         string
	Difficulty   string
	FunctionName string
	Parameters   []string
	Tests        []DBQuestion_Test
}
