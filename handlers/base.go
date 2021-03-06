package handlers

import (
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/validate", validateToken)
	mux.HandleFunc("/auth/logout", logout)

	// OauthGoogle
	mux.HandleFunc("/auth/google/login", oauthGoogleLogin)
	mux.HandleFunc("/auth/google/callback", oauthGoogleCallback)

	return mux
}
