package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"project-dpwims/database"
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

const urlRoutes = "/routes"
const urlRoutesId = "/routes/{id}"

const urlSchedules = "/schedules"
const urlSchedulesId = "/schedules/{id}"

const urlStopSchedules = "/stopschedules"
const urlStopSchedulesId = "/stopschedules/{id}"

const clientID = "train-service"
const keepAlive = 3 * time.Second
const connectRetryInterval = 3 * time.Second

// main, runs with this command in the terminal: docker compose --env-file ./env/develop.env up --build
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

	repositoryRoute := repositories.NewMySQLRepositoryRoute(db)
	routeService := services.NewRouteService(repositoryRoute)
	routeHandler := handlers.NewRouteHandler(routeService)

	repositorySchedule := repositories.NewMySQLScheduleRepository(db)
	scheduleService := services.NewScheduleService(repositorySchedule)
	scheduleHandler := handlers.NewScheduleHandler(scheduleService)

	repositoryStopSchedule := repositories.NewMySQLStopScheduleRepository(db)
	stopScheduleService := services.NewStopScheduleService(repositoryStopSchedule)
	stopScheduleHandler := handlers.NewStopScheduleHandler(stopScheduleService)

	router := chi.NewRouter()

	router.Post(urlTrains, trainHandler.CreateTrain)
	router.Get(urlTrains, trainHandler.GetAllTrains)
	router.Get(urlTrainsId, trainHandler.GetTrain)
	router.Delete(urlTrainsId, trainHandler.DeleteTrain)
	router.Patch(urlTrainsId, trainHandler.UpdateTrain)

	router.Post("/trains/{trainUUID}/arrived", trainHandler.MarkTrainArrived)

	router.Post(urlStations, stationHandler.CreateStation)
	router.Get(urlStations, stationHandler.GetAllStations)
	router.Get(urlStationsId, stationHandler.GetStation)
	router.Delete(urlStationsId, stationHandler.DeleteStation)
	router.Patch(urlStationsId, stationHandler.UpdateStation)

	router.Post(urlRoutes, routeHandler.CreateRoute)
	router.Get(urlRoutes, routeHandler.GetAllRoutes)
	router.Get(urlRoutesId, routeHandler.GetRoute)
	router.Delete(urlRoutesId, routeHandler.DeleteRoute)
	router.Patch(urlRoutesId, routeHandler.UpdateRoute)

	router.Post(urlSchedules, scheduleHandler.CreateSchedule)
	router.Get(urlSchedules, scheduleHandler.GetAllSchedules)
	router.Get(urlSchedulesId, scheduleHandler.GetSchedule)
	router.Delete(urlSchedulesId, scheduleHandler.DeleteSchedule)
	router.Patch(urlSchedulesId, scheduleHandler.UpdateSchedule)

	router.Post(urlStopSchedules, stopScheduleHandler.CreateStopSchedule)
	router.Get(urlStopSchedules, stopScheduleHandler.GetAllStopSchedules)
	router.Get(urlStopSchedulesId, stopScheduleHandler.GetStopSchedule)
	router.Delete(urlStopSchedulesId, stopScheduleHandler.DeleteStopSchedule)
	router.Patch(urlStopSchedulesId, stopScheduleHandler.UpdateStopSchedule)

	log.Println("Trains Service running on port 8082 with url http://localhost:8082")

	errorHttp := http.ListenAndServe(":8082", router)
	if errorHttp != nil {
		return
	}

}
