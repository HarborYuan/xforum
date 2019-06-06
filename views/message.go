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
