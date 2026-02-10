package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)

	fmt.Println(dsn)
	/*
		db, errorConnection := database.NewMySQLConnection(dsn)
		if errorConnection != nil {
			log.Fatal(errorConnection)
		}
		database.RunInitScripts(db)

		repository := repositories.NewMySQLRepositoryUsers(db)
		service := services.NewUserService(repository)
		handler := handlers.NewUserHandler(service)
		router := chi.NewRouter()

		router.Post(urlUsers, handler.CreateTrain)
		router.Get(urlUsers, handler.GetAllTrains)
		router.Get(urlUsersID, handler.GetTrain)
		router.Delete(urlUsersID, handler.DeleteUser)
		router.Patch(urlUsersID, handler.UpdateUser)
	*/
	log.Println("Trains Service running on port 8082 with url http://localhost:8082")
	/*
		errorHttp := http.ListenAndServe(":8082", router)
		if errorHttp != nil {
			return
		}
	*/
}
