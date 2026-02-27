package events

import (
	"encoding/json"
	"log"
	topics "project-dpwims/shared/mqtt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type TrainEvent struct {
	Event string `json:"event"`
	Time  string `json:"time"`
}

// PublishTrainEvent a function to publish a message for notifying the status of the train to the subscribers
func PublishTrainEvent(client mqtt.Client, trainUUID string, event TrainEvent) error {
	payload, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("Error marshaling train event: %v", err)
		return err
	}
	topic := topics.TrainEventsTopicFor(trainUUID)
	token := client.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}
