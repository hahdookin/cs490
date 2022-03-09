package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	L "github.com/AOrps/cs490/middleware/pyrun"
)

/*
Send POST request as `x-www-form-urlencoded`
*/

const (
	PORT    = 8087
	BACKEND = "https://web.njit.edu/~gmo9/back-end/backend.php"
)

type UP struct {
	username string `json:"username"`
	password string `json:"password"`
}

type Grade struct {
	Qid  string `json:"qid"`
	Code string `json:"code"`
}

// sendPOSTJSON -> sends a POST encoded with JSON
func sendPostJSON(endpoint string, cd UP) string {
	data := url.Values{
		"username": {cd.username},
		"password": {cd.password},
	}

	resp, err := http.PostForm(endpoint, data)
	L.Check(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	L.Check(err)

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

// func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(httpStatusCode)
// 	resp := make(map[string]string)
// 	resp["message"] = message
// 	jsonResp, _ := json.Marshal(resp)
// 	w.Write(jsonResp)
// }

func autograde(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		rBody, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		L.Check(err)
		var res []Grade

		// json.Unmarshal(rBody, &res)
		err = json.Unmarshal(rBody, &res)
		// decoder.DisallowUnknownFields()
		// err := decoder.Decode(&res)
		L.Check(err)

		// qid := r.FormValue("qid")
		// code := r.FormValue("code")

		// res.qid = qid
		// res.code = code
		for i, _ := range res {
			qid := res[i].Qid
			code := res[i].Code
			file := L.CreatePyFile(code, qid)
			output := L.RunCode(file)
			fmt.Fprintf(w, "%s: %s\n", qid, output)
		}

		// res := make(map[string]string)
		// res["qid"] = qid
		// res["code"] = code

		// enc := json.NewEncoder(w)
		// enc.Encode(res)
		// form values

		// fmt.Fprintf(w, "%s", string(marsh))
		// fmt.Println("ssss")

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
