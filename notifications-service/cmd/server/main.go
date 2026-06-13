package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"project-dpwims/database"

	"notifications-service/internal/api"
	"notifications-service/internal/mqtt"
	"notifications-service/internal/repository"
	"notifications-service/internal/service"
	sharedAuth "project-dpwims/shared/auth"
	topics "project-dpwims/shared/mqtt"

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

	subscriptionRepository := repository.NewMySQLSubscriptionRepository(db)

	mqttClient := mqtt.NewClient()
	if err := mqttClient.Connect(); err != nil {
		log.Fatal(err)
	}

	dispatcher := service.NewDispatcher(subscriptionRepository, mqttClient)
	log.Println("Connected to MQTT Broker at", os.Getenv("MQTT_BROKER_URL"))

	mqttClient.Subscribe(topics.TrainEventsTopic, 0, mqtt.TrainEventHandler(dispatcher))
	mqttClient.Subscribe(topics.TrainStopsTopic, 0, mqtt.TrainEventHandler(dispatcher))
	mqttClient.Subscribe(topics.TrainDelayTopic, 0, mqtt.TrainEventHandler(dispatcher))
	log.Println("MQTT client subscribed to topics")

	httpHandler := api.NewHandler(subscriptionRepository)
	router := api.NewRouter(httpHandler)

	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(sharedAuth.ValidateJWT)
	})

	log.Println("Subscription Service running on port 8084")
	if err := http.ListenAndServe(":8084", router); err != nil {
		log.Fatal(err)
	}
}
