package views

import (
	"encoding/json"
	"fmt"
	"github.com/HarborYuan/xforum/model/dbop"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	SnedMessage is a function that provides private message service
	{
		"sendee" : 2,
		"content" : "Hi Bob, I am fascinated by your face!"
	}
*/
// Haobo : Auto generated from https://mholt.github.io/json-to-go/
type JsonAddMessage struct {
	Sendee  int    `json:"sendee"`
	Content string `json:"content"`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var info JsonAddMessage
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
	if session.Values["loggedin"] == "false" {
		_, _ = w.Write([]byte("U200"))
		return
	}

	usernow := fmt.Sprint(session.Values["username"])
	useridnow := dbop.GetUid(usernow)
	if useridnow == -1 {
		_, _ = w.Write([]byte("M105"))
		return
	}
	if dbop.GetUserName(info.Sendee) == "@" {
		_, _ = w.Write([]byte("M105"))
		return
	}
	flag := dbop.AddMessage(
		useridnow,
		info.Sendee,
		info.Content)
	_, err = w.Write([]byte(flag))
}

/*
	getMessage is a function that handles get message requests
	{
		"sendee": 2
	}
*/

/*
return:
{
  "messages": [
    {
      "sender": "alice",
      "sendee": "bob",
      "createtime": "2019-06-06 11:43:38",
      "content": "Hi Bob, I am fascinated by your face!"
    },
    {
      "sender": "alice",
      "sendee": "bob",
      "createtime": "2019-06-06 11:54:22",
      "content": "Hi Bob, I am fascinated by your face!"
    },
    {
      "sender": "bob",
      "sendee": "alice",
      "createtime": "2019-06-06 13:54:58",
      "content": "Hi Alice, you're a good man!"
    }
  ]
}
*/

type JsonGetMessage struct {
	Sendee int `json:"sendee"`
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["loggedin"] != "true" {
		_, _ = w.Write([]byte("U200"))
		return
	}
	var info JsonGetMessage
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
	usernow := fmt.Sprint(session.Values["username"])
	useridnow := dbop.GetUid(usernow)
	if useridnow == -1 {
		_, _ = w.Write([]byte("G106"))
		return
	}
	if dbop.GetUserName(info.Sendee) == "@" {
		_, _ = w.Write([]byte("G106"))
		return
	}

	flag := dbop.GetMessage(useridnow, info.Sendee)
	_, err = w.Write([]byte(flag))
}
