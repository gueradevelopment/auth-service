package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"auth-service/handlers"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    addr,
		Handler: handlers.New(),
	}

	log.Printf("Starting Authentication HTTP Server. Listening at %q", os.Getenv("PORT"))
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
