package main

import (
	"log"
	"net/http"
	"user-service/internal/database"
	"user-service/internal/handlers"
	"user-service/internal/repositories"
	"user-service/internal/services"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

const pattern = "/users/{id}"

func main() {
	var dsn = "root:root@tcp(localhost:3306)/identity_users"
	db, errorConnection := database.NewMySQLConnection(dsn)
	if errorConnection != nil {
		log.Fatal(errorConnection)
	}
	database.RunInitScripts(db)
	repository := repositories.NewMySQLRepositoryUsers(db)
	service := services.NewUserService(repository)
	handler := handlers.NewUserHandler(service)
	router := chi.NewRouter()

	router.Post("/users", handler.CreateUser)
	router.Get("/users", handler.GetAllUsers)
	router.Get(pattern, handler.GetUser)
	router.Delete(pattern, handler.DeleteUser)

	log.Println("User Service running on port 8081 with url http://localhost:8081")
	errorHttp := http.ListenAndServe(":8081", router)
	if errorHttp != nil {
		return
	}
}
