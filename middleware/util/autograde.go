package util

import (
	"fmt"
	"os"
	"os/exec"
)

type Grade struct {
	Qid  string `json:"qid"`
	Code string `json:"code"`
}

type Ret struct {
	Qid   string
	Score string
}

func CreatePyFile(code string, filename string) string {
	// creates a python3 file

	// fmt.Println(os.Executable())
	file := fmt.Sprintf("%s.py", filename)
	f, err := os.Create(file)
	Check(err)

	_, err = f.Write([]byte("#!/usr/bin/python3\n"))
	Check(err)

	// Writes code into file
	_, err = f.Write([]byte(code))
	Check(err)

	f.Close()
	return file
}

// RunCode -> run a python file in golang
func RunCode(file string) string {
	cmd := exec.Command("python", file)
	out, err := cmd.Output()
	Check(err)

	return string(out)
}
