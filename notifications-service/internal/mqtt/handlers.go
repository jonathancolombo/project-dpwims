package mqtt

import (
	"notifications-service/internal/service"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TrainEventHandler(dispatcher *service.Dispatcher) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		topic := msg.Topic()
		parts := strings.Split(topic, "/")
		trainUUID := parts[1]

		dispatcher.HandleTrainEvent(trainUUID, msg.Payload())
	}
}
