package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	L "github.com/AOrps/cs490/middleware/util"
)

const (
	PORT    = 8087
	BACKEND = "https://web.njit.edu/~gmo9/back-end/backend.php"
)

// login -> handles login functionality (returns type: student | teacher)
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

func cringe(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err.Error())
		}

		qid := r.FormValue("qid")
		code := r.FormValue("code")

		fmt.Fprintf(w, "{\"qid\":\"%s\",\"code\":\"%s\"}", qid, code)
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
		var res L.Grade

		err = json.Unmarshal(rBody, &res)
		L.Check(err)

		// for i, _ := range res {
		// 	qid := res[i].Qid
		// 	code := res[i].Code

		// 	file := L.CreatePyFile(code, qid)
		// 	output := L.RunCode(file)
		// 	fmt.Fprintf(w, "%s: %s\n", qid, output)
		// }

		enc := json.NewEncoder(w)
		enc.Encode(res)

	default:
		fmt.Fprintf(w, "POST plz")
	}
}

func main() {
	port := strconv.Itoa(PORT)

	// handler
	http.HandleFunc("/", login)
	http.HandleFunc("/autograde", autograde)
	http.HandleFunc("/cringe", cringe)

	// Prints where it is on localhost
	fmt.Printf("http://localhost:%s\n", port)

	// start Server at port
	err := http.ListenAndServe(":"+port, nil)
	L.Check(err)

}
