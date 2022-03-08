package pyrun

import (
	"fmt"
	"log"
	"os/exec"
)

type Ret struct {
	questionID string
	score      string
}

func Hello() {
	fmt.Println("Hello!")
}

// RunCode -> run a python file in golang
func RunCode(file string) {
	cmd := exec.Command("python3", file)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(out))
}
