package handlers

import (
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/validate", validateToken)

	// OauthGoogle
	mux.HandleFunc("/auth/google/login", oauthGoogleLogin)
	mux.HandleFunc("/auth/google/callback", oauthGoogleCallback)

	// OauthGithub
	mux.HandleFunc("/auth/github/login", oauthGithubLogin)
	// mux.HandleFunc("/auth/github/callback", oauthGithubCallback)

	return mux
}
