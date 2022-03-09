package util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Question struct {
	Qid  string `json:"qid"`
	Code string `json:"code"`
}

type Ret struct {
	Output string
	Pass   bool
}

// CreatePyFile: creates a temp python file
func CreatePyFile(code string, filename string) string {
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

func detectOS() string {
	path := ""
	os := runtime.GOOS

	switch os {
	case "windows":
		path = "C:\\Users\\Andres\\AppData\\Local\\Microsoft\\WindowsApps\\python.exe"
		// path = "/c/Users/Andres/AppData/Local/Microsoft/WindowsApps/python3.exe"
	case "linux":
		path = "/usr/bin/python3"
	default:
		path = "/usr/bin/python"
	}

	fmt.Println(path)

	return path
}

// RunCode: run a python file in golang
func RunCode(file string) string {
	cmd := exec.Command(detectOS(), file)
	out, err := cmd.Output()
	Check(err)

	return string(out)
}
