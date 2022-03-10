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
	Runs        bool
	NameCorrect bool
	Output      []string
	Pass        []bool // Ordered by TestCase
}

// [
// 	{
// 		"tesy": true,
// 		"tset": "test"
// 	}
// ]

// testing strings

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

// AddTestCase: adds a test case to the end of the file
func AddTestCase(input, file string) string {
	f, err := os.ReadFile(file)
	Check(err)
	return string(f)
	// pyfunc := GetStringInBetween(string(f), "def ", `(`)
	// fmt.Println(pyfunc)

	// f, err = os.OpenFile(file, os.O_RDWR, 0644)
	// Check(err)
	// // for param in range params:
	// // 	f.Write([]byte("param, "))

	// f.Write([]byte(fmt.Sprintf("%s(", pyfunc)))

}

// RunCode: run a python file in golang
func RunCode(file string) string {
	cmd := exec.Command(detectOS(), file)
	out, err := cmd.Output()
	Check(err)

	return string(out)
}

// FullGrade: does all the autograde stuff
func FullGrade(w http.ResponseWriter, q Question) Ret {

	var doesRun bool

	// creates temp py file
	file := CreatePyFile(q.Code, q.Qid)
	// fmt.Fprintf(w, "%v", file)

	// f, err := os.OpenFile(file, os.O_RDWR, 0644)
	// Check(err)

	// defer f.Close()

	// // Add newline to temp file
	// f.Write([]byte("\n"))

	// Get TestCase
	// tc := "1"

	// for range in amount of parameters

	// out := AddTestCase(tc, file)

	f, err := os.ReadFile(file)
	Check(err)
	out := string(f)

	fmt.Printf("<%v>\n", out)
	// f, err = os.ReadFile(file)
	// Check(err)
	DBQuest := DBGetJSON("http://ec2-3-92-132-35.compute-1.amazonaws.com/questions/1")
	fmt.Println(DBQuest)

	pyfunc := GetStringInBetween(string(out), "def ", `(`)

	correctFuncName := pyfunc == DBQuest.FunctionName

	output := RunCode(file)

	if output != " " {
		doesRun = true
	}
	// fmt.Printf("%v: %v\n", pyfunc, output)
	RemovePyFile(file)

	// fmt.Fprintf(w, "%v", output)
	return Ret{
		Runs:        doesRun,
		NameCorrect: correctFuncName,
		Output: []string{
			"yes",
		},
		Pass: []bool{
			true,
		},
	}
}
