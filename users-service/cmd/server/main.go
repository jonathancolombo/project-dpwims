package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"project-dpwims/database"
	"users-service/internal/handlers"
	"users-service/internal/repositories"
	"users-service/internal/services"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

const urlUsers = "/users"
const urlUsersID = "/users/{id}"

// main, runs with this command in the terminal: docker compose --env-file ./env/develop.env up --build
func main() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)

	db, errorConnection := database.NewMySQLConnection(dsn)
	if errorConnection != nil {
		log.Fatal(errorConnection)
	}
	repository := repositories.NewMySQLRepositoryUsers(db)
	service := services.NewUserService(repository)
	handler := handlers.NewUserHandler(service)
	router := chi.NewRouter()

	router.Post(urlUsers, handler.CreateUser)
	router.Get(urlUsers, handler.GetAllUsers)
	router.Get(urlUsersID, handler.GetUser)
	router.Delete(urlUsersID, handler.DeleteUser)
	router.Patch(urlUsersID, handler.UpdateUser)

	log.Println("User Service running on port 8081 with url http://localhost:8081")
	errorHttp := http.ListenAndServe(":8081", router)
	if errorHttp != nil {
		return
	}
}
