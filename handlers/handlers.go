package handlers

import (
	"github.com/HarborYuan/xforum/views"
	"io/ioutil"
	"log"
	"net/http"
)

func Handler() {
	//Handle static files
	staticServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticServer))

	//Handle APIs
	http.HandleFunc("/api/signup/", views.Sinup)
	http.HandleFunc("/api/checkusername/", views.CheckUsername)
	http.HandleFunc("/api/checkemail/", views.CheckEmail)
	http.HandleFunc("/api/login/", views.Login)
	http.HandleFunc("/api/logout/", views.Logout)
	http.HandleFunc("/api/getposts/", views.GetPosts)
	http.HandleFunc("/api/getboards/", views.GetBoards)
	http.HandleFunc("/api/getresponse/", views.GetResponse)
	http.HandleFunc("/api/addboards/", views.AddBoards)
	http.HandleFunc("/api/addposts/", views.AddPost)
	http.HandleFunc("/api/addresponse/", views.AddResponse)

	//Handle favicon.ico
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/icon/favicon.ico")
	})

	//Handle Index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Handling / : " + "'" + r.URL.Path + "'")
		indexFile, err := ioutil.ReadFile("templates/index.html")
		if err != nil {
			log.Print("Handler() : What's wrong with index.html?")
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		_, err = w.Write(indexFile)
	})
}
