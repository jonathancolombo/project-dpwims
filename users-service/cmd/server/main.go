package main

import (
	"log"
	"net/http"

	"project-dpwims/database"
	"users-service/internal/handlers"
	"users-service/pkg/repositories"
	"users-service/pkg/services"

	sharedAuth "project-dpwims/shared/auth"
	util "project-dpwims/shared/utilities"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

const urlUsers = "/users"
const urlUsersID = "/users/{id}"

func main() {
	dsn := util.ConstructDSN()
	db, errorConnection := database.NewMySQLConnection(dsn)

	if errorConnection != nil {
		log.Fatal(errorConnection)
	}

	repository := repositories.NewMySQLRepositoryUsers(db)
	service := services.NewUserService(repository)
	handler := handlers.NewUserHandler(service)
	router := chi.NewRouter()

	router.Post(urlUsers, handler.CreateUser)

	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(sharedAuth.ValidateJWT)
		chiRouter.Use(sharedAuth.RequireRole("admin"))
		chiRouter.Get(urlUsers, handler.GetAllUsers)
	})

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
