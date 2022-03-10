package util

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

type Question struct {
	Qid  string `json:"qid"`
	Code string `json:"code"`
}

type Ret struct {
	Runs        bool     `json:"runs"`        // if python file runs
	NameCorrect bool     `json:"namecorrect"` // if function name is correct
	Output      []string `json:"output"`      // output[i: i in range(tests)]
	Pass        []bool   `json:"pass"`        // pass[i: i in range(tests)]
}

// CreatePyFile: creates a temp python file
func CreatePyFile(code, filename string) string {
	file := fmt.Sprintf("%s.py", filename)
	// creates a python3 file
	f, err := os.Create(file)
	Check(err)

	// Writes code into file
	_, err = f.Write([]byte(code))
	Check(err)

	f.Close()
	return file
}

func RemovePyFile(filename string) {
	err := os.Remove(filename)
	Check(err)
}

// detectOS: detects OS
func detectOS() string {
	path := ""
	os := runtime.GOOS

	switch os {
	case "windows":
		path = `C:\Program Files\Python310\python.exe`
	case "linux":
		path = "/usr/bin/python3"
	default:
		path = "/usr/bin/python"
	}

	return path
}

func RevertFile(file string, version []byte) {
}

// AddTestCase: adds a test case to the end of the file
func AddTestCase(file string, qTest DBQuestion_Test) {

	data, err := os.ReadFile(file)
	Check(err)
	pyfunc := GetStringInBetween(string(data), "def ", `(`)

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Check(err)

	defer f.Close()

	_, err = f.WriteString("\n")
	Check(err)

	_, err = f.WriteString(fmt.Sprintf("print(%s(", pyfunc))
	Check(err)

	for i, arg := range qTest.Arguments {
		if i < len(qTest.Arguments)-1 {
			_, err = f.WriteString(fmt.Sprintf("%s,", arg))
			Check(err)
		} else {
			_, err = f.WriteString(arg)
			Check(err)
		}
	}

	_, err = f.WriteString("), end=\"\")")
	Check(err)
}

// RunCode: run a python file in golang
func RunCode(file string) (string, bool) {
	doesRun := true
	cmd := exec.Command(detectOS(), file)
	out, err := cmd.Output()
	if err != nil {
		doesRun = false
	}

	return string(out), doesRun
}

// FullGrade: does all the autograde stuff
func FullGrade(w http.ResponseWriter, q Question) Ret {

	var doesRun, correctFuncName bool
	var Exec []string
	var Succeed []bool

	endpoint := fmt.Sprintf("http://ec2-3-92-132-35.compute-1.amazonaws.com/questions/%s", q.Qid)
	DBQuest := DBGetJSON(endpoint)

	// creates temp py file
	file := CreatePyFile(q.Code, q.Qid)

	// creates an 'anchor' so that file can be re-written back to this version
	f, err := os.ReadFile(file)
	Check(err)

	// iterates thru test cases, runs it and then reverts
	for _, test := range DBQuest.Tests {

		AddTestCase(file, test)
		output, doesRun := RunCode(file)
		// g, err := os.ReadFile(file)
		// Check(err)
		// fmt.Println(string(g))

		Exec = append(Exec, output)
		Succeed = append(Succeed, doesRun)

		// Reverts File back to what user submitted
		os.WriteFile(file, f, 0644)
	}

	// // Add newline to temp file
	// f.Write([]byte("\n"))

	// Get TestCase
	// tc := "1"

	// for range in amount of parameters

	// out := AddTestCase(tc, file)

	Check(err)
	out := string(f)
	fmt.Println(out)

	// for test := range DBQuest.Tests {

	// }

	// fmt.Printf("<%v>\n", out)
	// f, err = os.ReadFile(file)

	// Check(err)
	// fmt.Println(DBQuest)

	pyfunc := GetStringInBetween(string(out), "def ", `(`)
	correctFuncName = pyfunc == DBQuest.FunctionName

	// fmt.Printf("%v: %v\n", pyfunc, output)
	RemovePyFile(file)

	// fmt.Fprintf(w, "%v", output)
	return Ret{
		Runs:        doesRun,
		NameCorrect: correctFuncName,
		Output:      Exec,
		Pass:        Succeed,
	}
}
