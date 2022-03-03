package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	PORT    = 8080
	BACKEND = "https://web.njit.edu/~gmo9/back-end/backend.php"
)

type UP struct {
	username string `json:"username"`
	password string `json:"password"`
}

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// sendPOSTJSON -> sends a POST encoded with JSON
func sendPostJSON(url string, cd UP) string {
	// data := map[string]string{"username": cd.user, "password": cd.passwd}
	// json_data, err := json.Marshal(data)
	// json_data := []byte(`{
	// 	"username":"jane",
	// 	"password":"apple"
	// }`)
	json_data, err := json.Marshal(cd)
	check(err)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))
	check(err)

	defer resp.Body.Close()
	fmt.Println(resp.StatusCode, http.StatusCreated)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response: ", string(body))
	check(err)

	return string(body)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yes")
	// switch r.Method {
	// case "POST":
	// 	check(r.ParseForm())
	// 	// if err := r.ParseForm(); err != nil {
	// 	// log.Fatal(err.Error())
	// 	// }

	// 	username := r.FormValue("username")
	// 	password := r.FormValue("password")

	// 	client_data := UP{
	// 		user:   username,
	// 		passwd: password,
	// 	}

	// default:
	// 	fmt.Fprintf(w, "POST plz")
	// }

}

func main() {
	// port := strconv.Itoa(PORT)

	r := sendPostJSON(BACKEND, UP{
		username: "jane",
		password: "apple",
	})

	fmt.Println(r)

	// handler
	// http.HandleFunc("/", handler)

	// // actually get the server to
	// err := http.ListenAndServe(":"+port, nil)
	// check(err)

}
