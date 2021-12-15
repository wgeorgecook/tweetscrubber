package main

import "net/http"

func startOauthFlow(w http.ResponseWriter, r *http.Request) {
	logger.Info("Incoming request for oauth flow start")
}

func oauthCallback(w http.ResponseWriter, r *http.Request) {
	logger.Info("Incoming callback request")
}
