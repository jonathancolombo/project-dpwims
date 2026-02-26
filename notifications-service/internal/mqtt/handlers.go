package mqtt

import (
	"log"
	"notifications-service/internal/service"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// TrainEventHandler returns an MQTT message handler that processes incoming train events.
func TrainEventHandler(dispatcher *service.Dispatcher) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		topic := msg.Topic()
		parts := strings.Split(topic, "/")
		trainUUID := parts[1]
		log.Println("MQTT EVENT RECEIVED:", msg.Topic(), string(msg.Payload()))
		dispatcher.HandleTrainEvent(trainUUID, msg.Payload())
	}
}
