package service

import (
	"context"
	"fmt"
	"log"
	"notifications-service/internal/repository"
)

type Dispatcher struct {
	repo repository.SubscriptionRepository
}

func NewDispatcher(repo repository.SubscriptionRepository) *Dispatcher {
	return &Dispatcher{repo: repo}
}

func (d *Dispatcher) HandleTrainEvent(trainUUID string, payload []byte) {
	ctx := context.Background()

	users, err := d.repo.GetUsersByTrainUUID(ctx, trainUUID)
	if err != nil {
		log.Println("Error fetching subscribers:", err)
		return
	}

	for _, userID := range users {
		fmt.Printf("Notify user %d about train %s: %s\n",
			userID, trainUUID, string(payload))
	}
}
