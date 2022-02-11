package main

import (
	"codeChallenge/config"
	"net/http"
)

func main() {
	serverConfig := config.LoadConfig()
	server := &http.Server{
		Addr: serverConfig.Addr,
		Handler: serverConfig.Handler,
	}
	server.ListenAndServe()
}
