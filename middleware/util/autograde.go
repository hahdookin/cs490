package util

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

var (
	ENDPOINT = "http://ec2-3-92-132-35.compute-1.amazonaws.com"
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

// RemovePyFile: remove python file
func RemovePyFile(filename string) {
	err := os.Remove(filename)
	Check(err)
}

// detectOS: detects OS and then gets python path
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

// RevertFile: reverts file back to some pre-determined state
func RevertFile(file string, version []byte) {
}

// AddTestCase: adds a test case to the end of the file
func AddTestCase(file string, args []string) {

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

	for i, arg := range args {
		_, err := strconv.Atoi(arg)
		if err != nil {
			strVar := fmt.Sprintf("\"%s\"", arg)

			if i < len(args)-1 {
				_, err = f.WriteString(fmt.Sprintf("%s,", strVar))
				Check(err)
			} else {
				_, err = f.WriteString(strVar)
				Check(err)
			}
		} else {
			if i < len(args)-1 {
				_, err = f.WriteString(fmt.Sprintf("%s,", arg))
				Check(err)
			} else {
				_, err = f.WriteString(arg)
				Check(err)
			}
		}
	}

	_, err = f.WriteString("), end=\"\")")
	Check(err)
}

// RunCode: run a python file in golang
func RunCode(file, validate string) (string, bool) {
	cmd := exec.Command(detectOS(), file)
	out, err := cmd.Output()
	if err != nil {
		return "", false
	}

	return string(out), validate == string(out)
}

// FullGrade: does all the autograde stuff
func FullGrade(w http.ResponseWriter, q Question) Ret {

	var doesRun, correctFuncName bool
	var Exec []string
	var Succeed []bool
	var runOut bytes.Buffer

	endpoint := fmt.Sprintf("%s/questions/%s", ENDPOINT, q.Qid)
	DBQuest := DBGetJSON(endpoint)

	// creates temp py file
	file := CreatePyFile(q.Code, q.Qid)

	// creates an 'anchor' so that file can be re-written back to this version
	f, err := os.ReadFile(file)
	Check(err)

	// TODO: get error code
	cmd := exec.Command(detectOS(), file)
	cmd.Stdout = &runOut
	if err := cmd.Run(); err != nil {
		fmt.Printf("err: %v\n", err)
		doesRun = false
	} else {
		fmt.Printf("err: %v\n", err)
		doesRun = true
	}
	fmt.Printf("%q\n", runOut.String())

	if doesRun {
		// iterates thru test cases, runs it and then reverts
		for _, test := range DBQuest.Tests {

			AddTestCase(file, test.Arguments)
			validate := test.Output
			output, trySuccess := RunCode(file, validate)

			// fmt.Printf("output: %s", output)
			// fmt.Printf("trySuccess: %v", trySuccess)
			// fmt.Printf("output: %s")

			// To Print out stuff uncomment lines below
			// g, err := os.ReadFile(file)
			// Check(err)
			// fmt.Println(string(g))

			Exec = append(Exec, output)
			Succeed = append(Succeed, trySuccess)

			// Reverts File back to what user submitted
			os.WriteFile(file, f, 0644)
		}
	}
	Check(err)
	out := string(f)
	// fmt.Println(out)

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
