package pyrun

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Ret struct {
	Qid   string
	Score string
}

func Check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func CreatePyFile(code string, filename string) string {
	// creates a python3 file

	fmt.Println(os.Executable())
	file := fmt.Sprintf("%s.py", filename)
	f, err := os.Create(file)
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
