package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("🚆 Railway Ticketing System API starting...")

	http.HandleFunc("/health", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.WriteHeader(http.StatusOK)
		_, err := responseWriter.Write([]byte("OK"))
		if err != nil {
			return
		}
	})

	address := ":8080"
	log.Printf("Server running on http://localhost%s\n\n", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}
