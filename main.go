package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"./handlers"
	"github.com/joho/godotenv"
	// "github.com/gueradevelopment/auth-service/handlers"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":3000"),
		Handler: handlers.New(),
	}

	log.Printf("Starting Authentication HTTP Server. Listening at %q", os.Getenv("AUTH_PORT"))
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
