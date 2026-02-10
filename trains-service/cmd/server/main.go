package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"project-dpwims/database"
	"trains-service/internal/handlers"
	"trains-service/internal/repositories"
	"trains-service/internal/services"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

const urlTrains = "/trains"
const urlTrainsId = "/trains/{id}"

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

	repository := repositories.NewMySQLRepositoryTrains(db)
	service := services.NewTrainService(repository)
	handler := handlers.NewTrainHandler(service)
	router := chi.NewRouter()

	router.Post(urlTrains, handler.CreateTrain)
	router.Get(urlTrains, handler.GetAllTrains)
	router.Get(urlTrainsId, handler.GetTrain)
	router.Delete(urlTrainsId, handler.DeleteTrain)
	router.Patch(urlTrainsId, handler.UpdateTrain)

	log.Println("Trains Service running on port 8082 with url http://localhost:8082")

	errorHttp := http.ListenAndServe(":8082", router)
	if errorHttp != nil {
		return
	}

}
