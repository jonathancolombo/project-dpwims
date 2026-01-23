package main

import (
	"log"
	"net/http"
	"user-service/internal/handlers"
	"user-service/internal/repositories"
	"user-service/internal/services"

	"github.com/go-chi/chi/v5"
)

const pattern = "/users/{id}"

func main() {
	repository := repositories.NewInMemoryUserRepository()
	service := services.NewUserService(repository)
	handler := handlers.NewUserHandler(service)

	router := chi.NewRouter()

	router.Post("/users", handler.CreateUser)
	router.Get("/users", handler.GetAllUsers)
	router.Get(pattern, handler.GetUser)
	router.Delete(pattern, handler.DeleteUser)

	log.Println("User Service running on port 8081 with url http://localhost:8081")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		return
	}
}
