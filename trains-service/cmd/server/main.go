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
const urlTrainsId = "/trains/{uuid}"

const urlStations = "/stations"
const urlStationsId = "/stations/{id}"

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

	repositoryTrains := repositories.NewMySQLRepositoryTrains(db)
	trainService := services.NewTrainService(repositoryTrains)
	trainHandler := handlers.NewTrainHandler(trainService)

	repositoryStation := repositories.NewMySQLRepositoryStation(db)
	stationService := services.NewStationService(repositoryStation)
	stationHandler := handlers.NewStationHandler(stationService)

	router := chi.NewRouter()

	router.Post(urlTrains, trainHandler.CreateTrain)
	router.Get(urlTrains, trainHandler.GetAllTrains)
	router.Get(urlTrainsId, trainHandler.GetTrain)
	router.Delete(urlTrainsId, trainHandler.DeleteTrain)
	router.Patch(urlTrainsId, trainHandler.UpdateTrain)

	router.Post(urlStations, stationHandler.CreateStation)
	router.Get(urlStations, stationHandler.GetAllStations)
	router.Get(urlStationsId, stationHandler.GetStation)
	router.Delete(urlStationsId, stationHandler.DeleteStation)
	router.Patch(urlStationsId, stationHandler.UpdateStation)

	log.Println("Trains Service running on port 8082 with url http://localhost:8082")

	errorHttp := http.ListenAndServe(":8082", router)
	if errorHttp != nil {
		return
	}

}
