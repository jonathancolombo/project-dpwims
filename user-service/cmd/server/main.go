package main

import (
	"log"
	"net/http"
	"user-service/internal/handlers"
	"user-service/internal/repositories"
	"user-service/internal/services"
)

func main() {
	repo := repositories.NewInMemoryUserRepository()
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	http.HandleFunc("/users", handler.CreateUser)
	http.HandleFunc("/user", handler.GetUser)

	log.Println("User Service running on port 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		return
	}
}
