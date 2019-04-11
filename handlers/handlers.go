package handlers

import (
	"log"
	"net/http"
)

func Handler() {
	staticServer := http.FileServer(http.Dir("static"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("asd")
	})
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/icon/favicon.ico")
	})
	http.Handle("/static/", http.StripPrefix("/static/", staticServer))
}
