package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"project-dpwims/database"
	sharedAuth "project-dpwims/shared/auth"
	"time"
	"trains-service/internal/handlers"
	"trains-service/internal/repositories"
	"trains-service/internal/services"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

const urlTrains = "/trains"
const urlTrainsId = "/trains/{uuid}"

const urlStations = "/stations"
const urlStationsId = "/stations/{id}"

const urlSchedules = "/schedules"
const urlSchedulesId = "/schedules/{id}"

const urlStopSchedules = "/stopschedules"
const urlStopSchedulesId = "/stopschedules/{id}"

const clientID = "train-service"
const keepAlive = 3 * time.Second
const connectRetryInterval = 3 * time.Second

// main, runs with this command in the terminal: docker compose --env-file ./env/.env up --build
func main() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)
	brokerURL := os.Getenv("MQTT_BROKER")

	db, errorConnection := database.NewMySQLConnection(dsn)
	if errorConnection != nil {
		log.Fatal(errorConnection)
	}

	options := mqtt.NewClientOptions()
	options.AddBroker(brokerURL)
	options.SetClientID(clientID)
	options.SetCleanSession(true)
	options.SetAutoReconnect(true)
	options.SetConnectRetry(true)
	options.SetConnectRetryInterval(keepAlive)
	options.SetKeepAlive(connectRetryInterval)
	mqttClient := mqtt.NewClient(options)

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Mqtt connection failed %v", token.Error())
	}
	log.Printf("Connected to MQTT broker %s", brokerURL)

	repositoryTrains := repositories.NewMySQLRepositoryTrains(db)
	trainService := services.NewTrainService(repositoryTrains, mqttClient)
	trainHandler := handlers.NewTrainHandler(trainService)

	repositoryStation := repositories.NewMySQLRepositoryStation(db)
	stationService := services.NewStationService(repositoryStation)
	stationHandler := handlers.NewStationHandler(stationService)

	repositorySchedule := repositories.NewMySQLScheduleRepository(db)
	scheduleService := services.NewScheduleService(repositorySchedule)
	scheduleHandler := handlers.NewScheduleHandler(scheduleService)

	repositoryStopSchedule := repositories.NewMySQLStopScheduleRepository(db)
	stopScheduleService := services.NewStopScheduleService(repositoryStopSchedule)
	stopScheduleHandler := handlers.NewStopScheduleHandler(stopScheduleService)

	router := chi.NewRouter()
	// ROTTE PUBBLICHE
	router.Get(urlTrains, trainHandler.GetAllTrains)
	router.Get(urlTrainsId, trainHandler.GetTrain)
	router.Get(urlSchedules, scheduleHandler.GetAllSchedules)

	// ROTTE ADMIN
	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(sharedAuth.ValidateJWT)
		chiRouter.Use(sharedAuth.RequireRole("admin"))
		chiRouter.Post(urlTrains, trainHandler.CreateTrain)
		chiRouter.Patch(urlTrainsId, trainHandler.UpdateTrain)
		chiRouter.Delete(urlTrainsId, trainHandler.DeleteTrain)
		chiRouter.Post(urlStations, stationHandler.CreateStation)
		chiRouter.Get(urlStations, stationHandler.GetAllStations)
		chiRouter.Get(urlStationsId, stationHandler.GetStation)
		chiRouter.Delete(urlStationsId, stationHandler.DeleteStation)
		chiRouter.Patch(urlStationsId, stationHandler.UpdateStation)
		chiRouter.Post(urlSchedules, scheduleHandler.CreateSchedule)
		chiRouter.Get(urlSchedulesId, scheduleHandler.GetSchedule)
		chiRouter.Delete(urlSchedulesId, scheduleHandler.DeleteSchedule)
		chiRouter.Patch(urlSchedulesId, scheduleHandler.UpdateSchedule)
		chiRouter.Post(urlStopSchedules, stopScheduleHandler.CreateStopSchedule)
		chiRouter.Delete(urlStopSchedulesId, stopScheduleHandler.DeleteStopSchedule)
		chiRouter.Patch(urlStopSchedulesId, stopScheduleHandler.UpdateStopSchedule)
		chiRouter.Post("/trains/{trainUUID}/arrived", trainHandler.MarkTrainArrived)
		chiRouter.Get("/stopschedules/schedule/{scheduleId}", stopScheduleHandler.GetStopsBySchedule)
	})

	log.Println("Trains Service running on port 8082 with url http://localhost:8082")

	errorHttp := http.ListenAndServe(":8082", router)
	if errorHttp != nil {
		return
	}

}
