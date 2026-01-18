package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "develop"
	}
	err := godotenv.Load(".env/" + env + ".env")
	if err != nil {
		log.Printf("No .env file found for %s environment", env)
	}

	log.Println("🚆 Railway Ticketing System API starting...")

	http.HandleFunc("/health", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.WriteHeader(http.StatusOK)
		_, err := responseWriter.Write([]byte("OK"))
		if err != nil {
			return
		}
	})

	address := "0.0.0.0:8080"
	log.Printf("Server running on http://localhost%s\n\n", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}
