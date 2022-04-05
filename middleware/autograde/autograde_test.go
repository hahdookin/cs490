package autograde

import (
	"fmt"
	"os"
	"testing"
)

// Checks if file can be created and if it exists (then deletes file)
func TestCreatePyFile(t *testing.T) {
	trueFName := "test"
	fName := fmt.Sprintf("%s.py", trueFName)

	CreatePyFile(`def test():\n\tprint("hello")`, trueFName)
	_, err := os.Stat(fName)
	if err != nil {
		errMsg := fmt.Sprintf("file not found: %s could not be found", fName)
		t.Errorf(errMsg)
	}
	RemovePyFile(fName)
}
