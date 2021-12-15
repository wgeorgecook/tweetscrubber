package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func startHttpServer() {
	logger.Info("Start HTTP server")
	r := mux.NewRouter()
	r.HandleFunc("/oauth2/redirect", startOauthFlow)
	r.HandleFunc("/oauth2/callback", oauthCallback)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
