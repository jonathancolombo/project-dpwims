package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"project-dpwims/database"
	"users-service/internal/handlers"
	"users-service/pkg/repositories"
	"users-service/pkg/services"

	sharedAuth "project-dpwims/shared/auth"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

const urlUsers = "/users"
const urlUsersID = "/users/{id}"

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

	// ROTTE PUBBLICHE (chiunque puo' creare un utente)
	router.Post(urlUsers, handler.CreateUser)

	// -------------------------
	// ROTTE SOLO ADMIN
	// -------------------------
	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(sharedAuth.ValidateJWT)
		chiRouter.Use(sharedAuth.RequireRole("admin"))
		chiRouter.Get(urlUsers, handler.GetAllUsers)
	})

	// -------------------------
	// ROTTE SELF OR ADMIN
	// -------------------------
	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(sharedAuth.ValidateJWT)
		chiRouter.Use(sharedAuth.RequireSelfOrAdmin())

		chiRouter.Get(urlUsersID, handler.GetUser)
		chiRouter.Patch(urlUsersID, handler.UpdateUser)
		chiRouter.Delete(urlUsersID, handler.DeleteUser)
	})

	log.Println("User Service running on port 8081 with url http://localhost:8081")
	errorHttp := http.ListenAndServe(":8081", router)
	if errorHttp != nil {
		return
	}
}
