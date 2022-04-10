package autograde

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	U "github.com/AOrps/cs490/middleware/util"
)

var (
	ENDPOINT = "http://ec2-3-92-132-35.compute-1.amazonaws.com"
)

type Question struct {
	Qid        string `json:"qid"`
	Code       string `json:"code"`
	Constraint string `json:"constraint"`
}

type Ret struct {
	NameCorrect bool     `json:"namecorrect"` // if function name is correct
	Output      []string `json:"output"`      // output[i: i in range(tests)]
	Pass        []bool   `json:"pass"`        // pass[i: i in range(tests)]
	Constraint  bool     `json:"constraint"`
}

// CreatePyFile: creates a temp python file
func CreatePyFile(code, filename string) string {
	file := fmt.Sprintf("%s.py", filename)
	// creates a python3 file
	f, err := os.Create(file)
	U.Check(err)

	// Writes code into file
	_, err = f.Write([]byte(code))
	U.Check(err)

	f.Close()
	return file
}

// RemovePyFile: remove python file
func RemovePyFile(filename string) {
	err := os.Remove(filename)
	U.Check(err)
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

// CommentPreprocessing: eliminates single-line & multi-line comments
func CommentPreprocessing(original string) string {

	var res string

	// remove single line comments
	// start: #  &&  end: \n || end_of_line
	singleLine := regexp.MustCompile(`#.*?\n|#.*$`)
	removedSingleLineComment := singleLine.ReplaceAllString(original, "\n")

	// remove multi-line commments
	// start: """ && end: """ || start: ''' && end: '''
	multiLineDoubleQuotes := regexp.MustCompile(`""".*"""|'''.*'''`)
	removedCase1 := multiLineDoubleQuotes.ReplaceAllString(removedSingleLineComment, "")

	res = removedCase1

	return res
}

// findConstraint: find if user program (file) actually has a constraint
func findConstraint(fName, constraint string) bool {

	data, err := os.ReadFile(fName)
	U.Check(err)
	text := string(data)

	switch constraint {
	case "for":
		if strings.Contains(text, constraint) {
			return true
		}
	case "while":
		if strings.Contains(text, constraint) {
			return true
		}
	case "recursion":
		pyfunc, err := U.GetStringInBetween(text, "def ", `(`)
		if err != nil {
			return false
		} else {
			if strings.Contains(text, pyfunc) {
				return true
			}
		}
	default:
		return false
	}
	return false
}

// AddTestCase: adds a test case to the end of the file
func AddTestCase(file string, args []string) error {

	data, err := os.ReadFile(file)
	U.Check(err)

	//pyfunc : get's name of the user inputted function
	pyfunc, err := U.GetStringInBetween(string(data), "def ", `(`)
	if err != nil {
		return errors.New("can not parse out pyfunc")
	}

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	U.Check(err)

	defer f.Close()

	// add a new line to file
	_, err = f.WriteString("\n")
	U.Check(err)

	_, err = f.WriteString(fmt.Sprintf("print(%s(", pyfunc))
	U.Check(err)

	for i, arg := range args {
		_, err := strconv.Atoi(arg)
		if err != nil {
			strVar := fmt.Sprintf("\"%s\"", arg)

			if i < len(args)-1 {
				f.WriteString(fmt.Sprintf("%s,", strVar))
			} else {
				f.WriteString(strVar)
			}
		} else {
			if i < len(args)-1 {
				f.WriteString(fmt.Sprintf("%s,", arg))
			} else {
				f.WriteString(arg)
			}
		}
	}

	_, err = f.WriteString("), end=\"\")")
	U.Check(err)
	return nil
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

	var correctFuncName bool
	var Exec []string
	var Succeed []bool

	endpoint := fmt.Sprintf("%s/questions/%s", ENDPOINT, q.Qid)
	DBQuest := U.DBGetJSON(endpoint)

	// preprocesses code to remove both single line and multiline comments before file is created
	processedCode := CommentPreprocessing(q.Code)
	// fmt.Printf("pre:\n----------\n%s\n----------\n", q.Code)
	// fmt.Printf("post:\n----------\n%s\n----------\n", processedCode)
	// fmt.Fprintf(w, "%v\n", processedCode)

	// creates temp py file
	file := CreatePyFile(processedCode, q.Qid)

	// find constraint here
	passConstraint := findConstraint(file, q.Constraint)

	// creates an 'anchor' so that file can be re-written back to this version
	f, err := os.ReadFile(file)
	U.Check(err)
	// fmt.Printf("File looks like:\n-----\n%s\n-----\n", string(f))

	// iterates thru test cases, runs it and then reverts
	for _, test := range DBQuest.Tests {
		err = AddTestCase(file, test.Arguments)
		if err != nil {
			return Ret{}
		}

		validate := test.Output
		output, trySuccess := RunCode(file, validate)

		// To Print out stuff uncomment lines below
		// g, err := os.ReadFile(file)
		// U.Check(err)
		// fmt.Println(string(g))

		Exec = append(Exec, output)
		Succeed = append(Succeed, trySuccess)

		// Reverts File back to what user submitted
		os.WriteFile(file, f, 0644)
	}

	U.Check(err)
	out := string(f)
	// fmt.Println(out)

	pyfunc, err := U.GetStringInBetween(string(out), "def ", `(`)
	if err != nil {
		correctFuncName = false
	} else {
		correctFuncName = pyfunc == DBQuest.FunctionName
	}

	RemovePyFile(file)

	return Ret{
		NameCorrect: correctFuncName,
		Output:      Exec,
		Pass:        Succeed,
		Constraint:  passConstraint,
	}
}
