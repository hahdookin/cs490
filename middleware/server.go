package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	auto "github.com/AOrps/cs490/middleware/autograde"
	util "github.com/AOrps/cs490/middleware/util"
)

const (
	PORT    = 8080
	BACKEND = "https://web.njit.edu/~gmo9/back-end/backend.php"
)

func cringe(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err.Error())
		}

		qid := r.FormValue("qid")
		code := r.FormValue("code")
		constraint := r.FormValue("constraint")

		q := auto.Question{
			Qid:        qid,
			Code:       code,
			Constraint: constraint,
		}

		out := auto.FullGrade(w, q)

		enc := json.NewEncoder(w)
		enc.Encode(out)
	default:
		fmt.Fprintf(w, "POST plz")
	}
}

func autograde(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w, r)
	switch r.Method {
	case "POST":
		var question auto.Question

		rBody, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		util.Check(err)

		err = json.Unmarshal(rBody, &question)
		util.Check(err)

		out := auto.FullGrade(w, question)

		enc := json.NewEncoder(w)
		enc.Encode(out)

	default:
		fmt.Fprintf(w, "POST plz")
	}
}

func main() {
	port := strconv.Itoa(PORT)

	// handler
	http.HandleFunc("/autograde", autograde)
	http.HandleFunc("/cringe", cringe)

	// Prints where it is on localhost
	fmt.Printf("http://localhost:%s\n", port)

	// start Server at port
	err := http.ListenAndServe(":"+port, nil)
	util.Check(err)

}
