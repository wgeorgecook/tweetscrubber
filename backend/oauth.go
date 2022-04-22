package main

import (
	"fmt"
	"net/http"

	cv "github.com/nirasan/go-oauth-pkce-code-verifier"
)

const (
	TwitterAuthorizeBaseEndpoint = "https://twitter.com/i/oauth2/authorize"
)

type pkceCode struct {
	verifier   string
	challenger string
}

// generateCodeChallenge uses the go-oauth-pkce-code-verifier
// package to return the verifier and challenge generated from that
// package
func generateCodeChallenge() (*pkceCode, error) {
	// Create code_verifier
	v, err := cv.CreateCodeVerifier()
	if err != nil {
		return nil, err
	}
	codeVerifier := v.String()

	// Create code_challenge with S256 method
	codeChallenge := v.CodeChallengeS256()

	return &pkceCode{
		verifier:   codeVerifier,
		challenger: codeChallenge,
	}, nil
}

// startOauthFlow is the entrypoint to obtaining a Twitter access token
// on behalf of the user. It makes the inital request to Twitter
func startOauthFlow(w http.ResponseWriter, r *http.Request) {
	logger.Info("Incoming request for oauth flow start")

	// build authorize endpoint with secrets and parameters
	// https://twitter.com/i/oauth2/authorize?response_type=code&
	// client_id=M1M5R3BMVy13QmpScXkzTUt5OE46MTpjaQ&
	// redirect_uri=https://www.example.com&
	// scope=tweet.read%20users.read%20&state=state&
	// code_challenge=challenge&code_challenge_method=plain
	pkce, err := generateCodeChallenge()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	endpoint := TwitterAuthorizeBaseEndpoint +
		"response_type=code&" +
		fmt.Sprintf("client_id=%v&", config.TwitterClientID) +
		fmt.Sprintf("redirect_uri=%v&", config.TwitterRedirectURI) +
		"scope=tweet.write&" +
		fmt.Sprintf("state=%v&", pkce.verifier) +
		fmt.Sprintf("code_challenge=%v&", pkce.challenger) +
		"code_challenge_method=S256"

	// redirect the user to the endpoint constructed above
	http.Redirect(w, r, endpoint, http.StatusFound)
	return
}

// oauthCallback is the second leg in the oauth flow. If the user
// granted us access from the startOauthFlow redirect, it will receive
// a code on the request that we can use here to get a bearer token
func oauthCallback(w http.ResponseWriter, r *http.Request) {
	logger.Info("Incoming callback request")
}
