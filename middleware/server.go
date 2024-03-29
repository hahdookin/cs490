package main

import (
	"encoding/json"
	"fmt"
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

func formDataLegacy(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err.Error())
		}

		qid := r.FormValue("qid")
		code := r.FormValue("code")
		constraint := r.FormValue("constraint")

		// fmt.Println(constraint)
		// for key, value := range r.Form {
		// 	fmt.Printf("%v:%v\n", key, value)
		// }

		q := auto.Question{
			Qid:        qid,
			Code:       code,
			Constraint: constraint,
		}

		out := auto.FullGrade(q)

		enc := json.NewEncoder(w)
		enc.Encode(out)
	default:
		fmt.Fprintf(w, "POST plz")
	}
}

func autogradeGoEnableCors(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w, r)
	switch r.Method {
	case "POST":
		var question auto.Question

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&question)

		defer r.Body.Close()
		util.Check(err)

		out := auto.FullGrade(question)

		enc := json.NewEncoder(w)
		enc.Encode(out)

	default:
		fmt.Fprintf(w, "POST plz")
	}
}

func autograde(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var question auto.Question

		dec := json.NewDecoder(r.Body)

		defer r.Body.Close()

		err := dec.Decode(&question)
		util.Check(err)

		// fmt.Fprintf(w, "%v", question)
		out := auto.FullGrade(question)

		enc := json.NewEncoder(w)
		enc.Encode(out)

	default:
		fmt.Fprintf(w, "POST plz")
	}
}

func main() {
	port := strconv.Itoa(PORT)

	// handler
	http.HandleFunc("/cringe", formDataLegacy)
	http.HandleFunc("/autograde", autogradeGoEnableCors)
	http.HandleFunc("/actual", autograde)

	// Prints where it is on localhost
	fmt.Printf("http://localhost:%s\n", port)

	// start Server at port
	err := http.ListenAndServe(":"+port, nil)
	util.Check(err)

}
