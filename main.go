package main

import (
	"github.com/HarborYuan/xforum/config"
	"github.com/HarborYuan/xforum/handlers"
	"log"
	"net/http"
)

func main() {
	err := config.LoadConfig("config.json")
	if err != nil {
		log.Print("main() : Failed to get the config info, exit!")
		return
	}
	listeningAddr := config.Config.Listen.ServerAddr + ":" + config.Config.Listen.ServerPort
	handlers.Handler()
	log.Print("Running server on " + listeningAddr)
	log.Fatal(http.ListenAndServe(listeningAddr, nil))
}
