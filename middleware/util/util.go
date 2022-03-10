package util

import (
	"log"
	"strings"
)

// Check: error handling
func Check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// GetStringInBetween: get substrinf from 2 (potentially) diffeent delimiters
func GetStringInBetween(str, start, end string) string {
	s := strings.Index(str, start)
	if s == -1 {
		log.Fatal("start of string is failing")
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		log.Fatal("end of string is failing")
	}
	return str[s : s+e]

}
