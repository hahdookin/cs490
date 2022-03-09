package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	L "github.com/AOrps/cs490/middleware/pyrun"
)

const (
	PORT    = 8087
	BACKEND = "https://web.njit.edu/~gmo9/back-end/backend.php"
)

// login -> server logic
func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		rBody, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		L.Check(err)
		var clientData L.UP

		err = json.Unmarshal(rBody, &clientData)
		L.Check(err)

		resp := L.SendPostJSON(BACKEND, clientData)
		fmt.Fprintf(w, "%v", resp)

	default:
		fmt.Fprintf(w, "POST plz")
	}
}

func autograde(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		rBody, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		L.Check(err)
		var res []L.Grade

		err = json.Unmarshal(rBody, &res)
		L.Check(err)

		for i, _ := range res {
			qid := res[i].Qid
			code := res[i].Code
			fmt.Fprintf(w, "%s: %s\n", qid, code)

			// file := L.CreatePyFile(code, qid)
			// output := L.RunCode(file)
			// fmt.Fprintf(w, "%s: %s\n", qid, output)
		}

	default:
		fmt.Fprintf(w, "POST plz")
	}
}

func main() {
	port := strconv.Itoa(PORT)

	// handler
	http.HandleFunc("/", login)
	http.HandleFunc("/autograde", autograde)

	// Prints where it is on localhost
	fmt.Printf("http://localhost:%s\n", port)

	// start Server at port
	err := http.ListenAndServe(":"+port, nil)
	L.Check(err)

}
