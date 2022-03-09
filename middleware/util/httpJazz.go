package util

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type UP struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SendPOSTJSON -> sends a POST encoded with JSON (frontend)
func SendPostJSON(endpoint string, cd UP) string {
	data := url.Values{
		"username": {cd.Username},
		"password": {cd.Password},
	}

	resp, err := http.PostForm(endpoint, data)
	Check(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)

	return string(body)
}
