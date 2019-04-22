package views

import (
	"encoding/json"
	"github.com/HarborYuan/xforum/model/dbop"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	Signup is a function that handles users' register requests.
	{
		"email" : "example@example.com",
		"username" : "Bob",
		"password" : "password",
		"gender" : 0,
		"birthday" : "2001-01-01"
	}
*/
// Haobo : Auto generated from https://mholt.github.io/json-to-go/
type JsonSignup struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   int    `json:"gender"`
	Birthday string `json:"birthday"`
}

func Sinup(w http.ResponseWriter, r *http.Request) {
	var info JsonSignup
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Unmarshal
	err = json.Unmarshal(b, &info)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Print(err)
		return
	}
	flag := dbop.AddUser(
		info.Username,
		info.Username,
		info.Password,
		info.Email,
		"",
		info.Birthday,
		string(info.Gender))

	if flag != true {
		log.Print("Error!")
	}
	_, err = w.Write([]byte("Success"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Print(err)
		return
	}
}

/*
	Login is a function that handles users' login requests.
	{
		"username" : "Bob",
		"password" : "password"
	}
*/
// Haobo : Auto generated from https://mholt.github.io/json-to-go/
type JsonLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var info JsonLogin
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Unmarshal
	err = json.Unmarshal(b, &info)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Print(err)
		return
	}
}
