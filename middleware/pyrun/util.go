package pyrun

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

type UP struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Grade struct {
	Qid  string `json:"qid"`
	Code string `json:"code"`
}

type Ret struct {
	Qid   string
	Score string
}

// SendPOSTJSON -> sends a POST encoded with JSON
func SendPostJSON(endpoint string, cd UP) string {
	data := url.Values{
		"username": {cd.Username},
		"password": {cd.Password},
	}

	resp, err := http.PostForm(endpoint, data)
	Check(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)

	return string(body)
}

func Check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
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
