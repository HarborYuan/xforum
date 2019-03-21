package main

import (
	"log"
	"net/http"
)

func main() {
	listeningAddr := "127.0.0.1:1024"
	log.Print("Running server on " + listeningAddr)
	staticServer := http.FileServer(http.Dir("static"))
	http.HandleFunc("/", CompleteTaskFunc)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/icon/favicon.ico")
	})
	http.Handle("/static/", http.StripPrefix("/static/", staticServer))
	log.Fatal(http.ListenAndServe(listeningAddr, nil))
}

func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(r.URL.Path))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Path : " + r.URL.Path)
}
