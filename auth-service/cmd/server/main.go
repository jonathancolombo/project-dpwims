package main

import (
	"log"
	"net/http"
	util "project-dpwims/shared/utilities"

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
	dsn := util.ConstructDSN()

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
