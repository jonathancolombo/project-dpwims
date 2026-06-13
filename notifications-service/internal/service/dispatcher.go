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

func (dispatcher *Dispatcher) HandleTrainEvent(scheduleID int64, payload []byte) {
	log.Println("Handling event for schedule", scheduleID)
	users, err := dispatcher.repository.GetUsersByScheduleID(context.Background(), scheduleID)
	if err != nil {
		log.Println("Error fetching subscribers:", err)
		return
	}
	log.Println("Got subscribers", len(users))
	for _, userID := range users {
		topic := fmt.Sprintf("notifications/user/%d", userID)
		if err := dispatcher.mqttClient.Publish(topic, 0, false, payload); err != nil {
			log.Println("Error publishing notification:", err)
		} else {
			log.Printf("Notification sent to user %d on topic %s", userID, topic)
		}
	}
}
