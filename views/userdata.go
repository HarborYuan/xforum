package views

import (
	"encoding/json"
	"github.com/HarborYuan/xforum/model/dbop"
	"github.com/gorilla/sessions"
	"io/ioutil"
	"log"
	"net/http"
)

// Gorilla session initiate
var Store *sessions.CookieStore

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
		info.Password,
		info.Email,
		info.Birthday,
		string(info.Gender))
	_, err = w.Write([]byte(flag))
}

/*
	CheckUsername is a function that handles
	{
		"username" : "Bob"
	}
*/
// Haobo : Auto generated from https://mholt.github.io/json-to-go/
type JsonCheckUsername struct {
	Username string `json:"username"`
}

func CheckUsername(w http.ResponseWriter, r *http.Request) {
	var info JsonCheckUsername
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
	flag := dbop.CheckUsername(info.Username)
	_, err = w.Write([]byte(flag))
}

/*
	CheckEmail is a function that handles
	{
		"email" : "example@example.com"
	}
*/
// Haobo : Auto generated from https://mholt.github.io/json-to-go/
type JsonCheckEmail struct {
	Email string `json:"email"`
}

func CheckEmail(w http.ResponseWriter, r *http.Request) {
	var info JsonCheckEmail
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
	flag := dbop.CheckEmail(info.Email)
	_, err = w.Write([]byte(flag))
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
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["loggedin"] == "true" {
		_, _ = w.Write([]byte("U200"))
		return
	}
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
	flag := dbop.CheckPass(info.Username, info.Password)
	if flag == "C100-1" {
		session.Values["loggedin"] = "true"
		session.Values["username"] = info.Username
		_ = session.Save(r, w)
	}
	_, err = w.Write([]byte(flag))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["loggedin"] != "true" {
		_, _ = w.Write([]byte("U200"))
		return
	}
	session.Values["loggedin"] = "false"
	session.Values["username"] = ""
	_ = session.Save(r, w)
	_, err = w.Write([]byte("U100"))
}

/*
	GetUserInfo is a function that xxxxxxxxxxxxxxxxxx
	{
		"uid" : 1
	}
*/
// Haobo : Auto generated from https://mholt.github.io/json-to-go/

type JsonGetUserInfo struct {
	UID int `json:"uid"`
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["loggedin"] != "true" {
		_, _ = w.Write([]byte("U200"))
		return
	}
	var info JsonGetUserInfo
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
	flag := dbop.GetUserInfo(info.UID)
	_, err = w.Write([]byte(flag))
}
