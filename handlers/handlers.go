package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Handler() {
	staticServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticServer))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/icon/favicon.ico")
	})
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
