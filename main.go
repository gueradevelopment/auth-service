package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"auth-service/handlers"
)

func main() {

	server := &http.Server{
		Addr:    fmt.Sprintf(":3000"),
		Handler: handlers.New(),
	}

	log.Printf("Starting Authentication HTTP Server. Listening at %q", os.Getenv("PORT"))
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
