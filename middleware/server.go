package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	PORT = 8080
)

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	port := strconv.Itoa(PORT)

	// handler
	http.HandleFunc("/", handler)

	// actually get the server to
	err := http.ListenAndServe(":"+port, nil)
	check(err)

}
