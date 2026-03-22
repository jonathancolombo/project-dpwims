package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"auth-service/internal/handlers"
	"auth-service/internal/routes"
	authServices "auth-service/internal/services"

	"project-dpwims/database"

	userRepositories "users-service/pkg/repositories"
	userServices "users-service/pkg/services"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)

	db, err := database.NewMySQLConnection(dsn)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := userRepositories.NewMySQLRepositoryUsers(db)
	userService := userServices.NewUserService(userRepo)

	authService := authServices.NewAuthService(userService)
	authHandler := handlers.NewAuthHandler(authService)

	router := chi.NewRouter()

	routes.RegisterAuthRoutes(router, authHandler)

	log.Println("Auth Service running on port 8082 with url http://localhost:8085")
	if err := http.ListenAndServe(":8085", router); err != nil {
		log.Fatal(err)
	}
}
