package service

import (
	"context"
	"fmt"
	"log"
	"notifications-service/internal/ports"
	"notifications-service/internal/repository"
)

type Dispatcher struct {
	repository repository.SubscriptionRepository
	mqttClient ports.MqttPublisher
}

func NewDispatcher(repo repository.SubscriptionRepository, publisher ports.MqttPublisher) *Dispatcher {
	return &Dispatcher{
		repository: repo,
		mqttClient: publisher,
	}
}

// HandleTrainEvent processes a train event by fetching the list of subscribed users and notifying them.
func (dispatcher *Dispatcher) HandleTrainEvent(trainUUID string, payload []byte) {
	log.Println("Handling train event", trainUUID)
	users, err := dispatcher.repository.GetUsersByTrainUUID(context.Background(), trainUUID)
	if err != nil {
		log.Println("Error fetching subscribers:", err)
		return
	}
	log.Println("Got subscribers", len(users))
	for _, userID := range users {
		topic := fmt.Sprintf("notifications/user/%d", userID)

		err := dispatcher.mqttClient.Publish(topic, 0, false, payload)
		if err != nil {
			log.Println("Error publishing notification:", err)
		} else {
			log.Printf("Notification sent to user %d on topic %s", userID, topic)
		}
	}

}
