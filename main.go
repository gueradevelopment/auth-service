package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/subosito/gotenv"

	"github.com/gueradevelopment/auth-service/handlers"
)

func init() {
	gotenv.Load()
}

func main() {

	server := &http.Server{
		Addr:    fmt.Sprintf(":3000"),
		Handler: handlers.New(),
	}

	log.Printf("Starting Authentication HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
