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

func Check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func CreatePyFile(code string) {

	// f, err := os.Create("data.py")
	// Check(err)
	//

}

// RunCode -> run a python file in golang
func RunCode(file string) {
	cmd := exec.Command("python3", file)
	out, err := cmd.Output()
	Check(err)

	fmt.Println(string(out))
}
