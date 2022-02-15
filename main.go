package main

import (
	"codeChallenge/config"
	"log"
	"net/http"
)

func main() {
	serverConfig := config.LoadConfig()
	server := &http.Server{
		Addr: serverConfig.Addr,
		Handler: serverConfig.Handler,
	}
	log.Fatal(server.ListenAndServe())
}
