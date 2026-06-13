package mqtt

import (
	"log"
	"notifications-service/internal/service"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TrainEventHandler(dispatcher *service.Dispatcher) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		topic := msg.Topic()
		parts := strings.Split(topic, "/")
		log.Println("MQTT EVENT RECEIVED:", topic, string(msg.Payload()))

		// trains/{uuid}/schedules/{scheduleID}/delay|stops|events
		if len(parts) < 4 {
			log.Println("Unexpected topic format:", topic)
			return
		}

		scheduleIDStr := parts[3]
		scheduleID, err := strconv.ParseInt(scheduleIDStr, 10, 64)
		if err != nil {
			log.Println("Failed to parse scheduleID from topic:", err)
			return
		}

		dispatcher.HandleTrainEvent(scheduleID, msg.Payload())
	}
}
