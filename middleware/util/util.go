package util

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Check: error handling
func Check(e error) {
	if e != nil {
		log.Fatal(e.Error())
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

// EnableCors: headers needed to enable cors
func EnableCors(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS,PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Referrer-Policy", "no-referrer")
}

func JustStartScreaming(w http.ResponseWriter) {
	fmt.Fprintf(w, "%s", "bruh")
}
