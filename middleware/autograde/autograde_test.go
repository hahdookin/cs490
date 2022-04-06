package autograde

import (
	"fmt"
	"os"
	"regexp"
	"testing"
)

// Checks if file can be created and if it exists (then deletes file)
func TestCreatePyFile(t *testing.T) {
	trueFName := "test"
	fName := fmt.Sprintf("%s.py", trueFName)
	code := "def test(x):\n\treturn x"

	CreatePyFile(code, trueFName)
	_, err := os.Stat(fName)
	if err != nil {
		errMsg := fmt.Sprintf("file not found: %s could not be found", fName)
		t.Errorf(errMsg)
	}
	RemovePyFile(fName)
}

func TestCommentPreprocessing(t *testing.T) {
	code := "def test(x):\n\t\"\"\"yo\"\"\"\n# no\n\treturn x"

	singleLine := regexp.MustCompile(`#.*\n`)
	removedSingleLineComment := singleLine.ReplaceAllString(code, "")

	// remove multi-line commments
	// start: """ &&& end: """
	multLine := regexp.MustCompile(`""".*"""`)
	res := multLine.ReplaceAllString(removedSingleLineComment, "")

	fmt.Println(res)
}
