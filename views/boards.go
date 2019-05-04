package views

import (
	"encoding/json"
	"github.com/HarborYuan/xforum/model/dbop"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	getPosts is a function that handles get posts requests
	{
		"path": "index"
	}
*/
/*
	return :
	{"posts":[{"uid":1,"createtime":"2019-03-02 20:29:57","content":"是不是太差劲了"},{"uid":2,"createtime":"2019-04-30 20:29:57","content":"差评"}]}
*/
type JsonGetPosts struct {
	Path string `json:"path"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["loggedin"] != "true" {
		_, _ = w.Write([]byte("U200"))
		return
	}
	var info JsonGetPosts
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
	flag := dbop.GetPosts(info.Path)
	_, err = w.Write([]byte(flag))
}
