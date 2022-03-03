package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yes")
	switch r.Method {
	case "POST":
		check(r.ParseForm())
		// if err := r.ParseForm(); err != nil {
		// log.Fatal(err.Error())
		// }

		user := r.FormValue("username")
		passwd := r.FormValue("password")

		clientData := UP{
			username: user,
			password: passwd,
		}

		r := sendPostJSON(BACKEND, clientData)

	default:
		fmt.Fprintf(w, "POST plz")
	}

}

func main() {
	port := strconv.Itoa(PORT)

	r := sendPostJSON(BACKEND, UP{
		username: "jane",
		password: "apple",
	})

	fmt.Println(r)

	// reqBody, err := json.Marshal(map[string]string{
	// 	"username": "jane",
	// 	"password": "apple",
	// })
	// check(err)
	// cd := UP{
	// 	username: "jane",
	// 	password: "apple",
	// }

	// data := url.Values{
	// 	"username": {cd.username},
	// 	"password": {cd.password},
	// }

	// resp, err := http.PostForm("https://web.njit.edu/~gmo9/back-end/backend.php", data)
	// check(err)

	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// check(err)

	// log.Println(string(body))

	// fmt.Println(r)

	// handler
	http.HandleFunc("/", handler)

	// start Server at port
	err := http.ListenAndServe(":"+port, nil)
	check(err)

}
