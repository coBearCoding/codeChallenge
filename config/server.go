package config

import (
	"codeChallenge/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Addr    string
	Handler http.Handler
}

func LoadConfig() Server {
	router := mux.NewRouter()

	router.HandleFunc("/api", handlers.Index).Methods(http.MethodPost, http.MethodOptions)
	router.Use(mux.CORSMethodMiddleware(router))

	serverConfig := Server{
		Addr:    ":3030",
		Handler: router,
	}

	return serverConfig
}
