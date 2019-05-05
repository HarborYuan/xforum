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

/*
	getBoards is a function that handles get boards
*/
// return {"board":[{"name":"科技","path":"tech"},{"name":"体育","path":"pe"},{"name":"教育","path":"edu"}]}
func GetBoards(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["loggedin"] != "true" {
		_, _ = w.Write([]byte("U200"))
		return
	}
	res := dbop.GetBoards()
	_, err = w.Write([]byte(res))
}

/*
	getBoards is a function that handles get boards
	{
		"pid": 1
	}
*/
// return : {"response":[{"uid":3,"createtime":"2019-05-04 19:44:01","content":"附议"},{"uid":4,"createtime":"2019-05-04 19:44:01","content":"反对"}]}
type JsonGetResponse struct {
	Pid int `json:"pid"`
}

func GetResponse(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["loggedin"] != "true" {
		_, _ = w.Write([]byte("U200"))
		return
	}

	var info JsonGetResponse
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
	flag := dbop.GetResponse(info.Pid)
	_, err = w.Write([]byte(flag))
}

/*
	{
		"name" : "体育",
		"path" : "pe"
	}
*/
// ! ---- Admin required  ---- !
// ########## must be modified when deploy
type JsonAddBoards struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func AddBoards(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["loggedin"] != "true" /*|| session.Values["isAdmin"] != "true" */ {
		_, _ = w.Write([]byte("U200"))
		return
	}
	var info JsonAddBoards
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
	flag := dbop.AddBoard(info.Name, info.Path)
	_, err = w.Write([]byte(flag))
}

/*
	{
		"path" : "pe",
		"content" : "我今天刷题了，求夸"
	}
*/
type JsonAddPost struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["loggedin"] != "true" {
		_, _ = w.Write([]byte("U200"))
		return
	}
	var info JsonAddPost
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
	username := session.Values["username"]
	strusername := fmt.Sprint(username)
	userid := dbop.GetUid(strusername)
	if userid == -1 {
		_, err = w.Write([]byte("B105"))
		return
	}
	flag := dbop.AddPost(info.Content, info.Path, userid)
	_, err = w.Write([]byte(flag))
}
