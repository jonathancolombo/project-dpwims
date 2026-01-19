package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	env := os.Getenv("APP_ENV")
	if env == "" {
		log.Println("APP_ENV env variable not set")
		env = "develop"
	}
	log.Println("APP_ENV env:", env)
	log.Println("🚆 Railway Ticketing System API starting...")

	http.HandleFunc("/health", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.WriteHeader(http.StatusOK)
		_, err := responseWriter.Write([]byte("OK"))
		if err != nil {
			return
		}
	})

	address := "0.0.0.0:8080"
	log.Printf("Server running on http://%s\n\n", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}
