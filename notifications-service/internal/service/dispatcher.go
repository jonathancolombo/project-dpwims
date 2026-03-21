package service

import (
	"context"
	"log"
	"notifications-service/internal/repository"
)

// Dispatcher is responsible for handling train events and notifying subscribed users.
type Dispatcher struct {
	repository repository.SubscriptionRepository
}

// NewDispatcher creates a new Dispatcher instance with the provided SubscriptionRepository.
func NewDispatcher(subscriptionRepository repository.SubscriptionRepository) *Dispatcher {
	return &Dispatcher{repository: subscriptionRepository}
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
		log.Println("Adding user", userID)
		log.Printf("Notify user %dispatcher about train %s: %s\n",
			userID, trainUUID, string(payload))
	}
}
