package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"notifications-service/internal/api"
	"notifications-service/internal/mqtt"
	"notifications-service/internal/repository"
	"notifications-service/internal/service"
	"project-dpwims/database"

	_ "github.com/go-sql-driver/mysql"
)

// Main function to start the notification service
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

	repo := repository.NewMySQLSubscriptionRepository(db)
	dispatcher := service.NewDispatcher(repo)

	mqttClient := mqtt.NewClient()
	if err := mqttClient.Connect(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MQTT Broker at ", os.Getenv("MQTT_BROKER_URL"))
	mqttClient.Subscribe(mqtt.TrainEventsTopic, 0, mqtt.TrainEventHandler(dispatcher))
	mqttClient.Subscribe(mqtt.TrainStopsTopic, 0, mqtt.TrainEventHandler(dispatcher))
	mqttClient.Subscribe(mqtt.TrainDelayTopic, 0, mqtt.TrainEventHandler(dispatcher))
	log.Println("Mqtt client subscribed to topics")

	httpHandler := api.NewHandler(repo)
	router := api.NewRouter(httpHandler)

	log.Println("Notification Service running on port 8084 with url http://localhost:8084")
	errorHttp := http.ListenAndServe(":8084", router)
	if errorHttp != nil {
		return
	}

}
