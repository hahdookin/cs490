package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

/*
Send POST request as `x-www-form-urlencoded`
*/

const (
	PORT    = 8087
	BACKEND = "https://web.njit.edu/~gmo9/back-end/backend.php"
)

type UP struct {
	username string
	password string
}

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// sendPOSTJSON -> sends a POST encoded with JSON
func sendPostJSON(endpoint string, cd UP) string {
	data := url.Values{
		"username": {cd.username},
		"password": {cd.password},
	}

	resp, err := http.PostForm(endpoint, data)
	check(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return string(body)
}

// login -> server logic
func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err.Error())
		}

		user := r.FormValue("username")
		passwd := r.FormValue("password")

		clientData := UP{
			username: user,
			password: passwd,
		}
		// fmt.Println(clientData)
		// fmt.Println(r.PostForm)

		// log.Printf("POST: user: [%s]  passwd: [%s]", user, passwd)

		resp := sendPostJSON(BACKEND, clientData)
		fmt.Fprintf(w, "%v", resp)

	default:
		fmt.Fprintf(w, "POST plz")
	}
}

func autograde(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err.Error())
		}

		qid := r.FormValue("qid")
		code := r.FormValue("code")

		fmt.Fprintf(w, "{qid:%s,code:%s}", qid, code)
		// form values

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
	check(err)

}
