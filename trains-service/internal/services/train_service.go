package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	topics "project-dpwims/shared/mqtt"
	"time"
	"trains-service/internal/events"
	"trains-service/internal/models"
	"trains-service/internal/repositories"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

// TrainService defines the interface for managing Train entities.
type TrainService struct {
	repository repositories.ITrainRepository
	mqttClient mqtt.Client
}

// NewTrainService creates a new TrainService instance
func NewTrainService(repository repositories.ITrainRepository, client mqtt.Client) *TrainService {
	return &TrainService{
		repository: repository,
		mqttClient: client,
	}
}

// CreateTrain creates a new train
func (trainService *TrainService) CreateTrain(context context.Context, train *models.Train) (*models.Train, error) {
	if train == nil {
		return nil, errors.New("train is nil")
	}

	if train.Number == "" {
		return nil, errors.New("train number is empty")
	}

	if train.Capacity <= 0 {
		fmt.Println("train capacity is less than or equal to 0, setting default capacity to 500")
		train.Capacity = 500
	}

	if train.Type == "" {
		fmt.Println("train type is empty, setting default type to regional")
		train.Type = models.TrainTypeRegional
	}

	if train.Status == "" {
		fmt.Println("train status is empty, setting default status to active")
		train.Status = models.StatusActive
	}

	train.UUID = uuid.NewString()
	return trainService.repository.Create(context, train)
}

// GetTrain retrieves a train by their UUID
func (trainService *TrainService) GetTrain(context context.Context, uuid string) (*models.Train, error) {
	if uuid == "" {
		return nil, errors.New("uuid must be different than empty")
	}
	return trainService.repository.GetByID(context, uuid)
}

// GetAllTrains retrieves all trains
func (trainService *TrainService) GetAllTrains(context context.Context) ([]*models.Train, error) {
	if trainService.repository == nil {
		return nil, errors.New("repository is nil")
	}
	return trainService.repository.GetAll(context)
}

// DeleteTrainByID deletes a train by their UUID
func (trainService *TrainService) DeleteTrainByID(context context.Context, uuid string) error {
	if uuid == "" {
		return errors.New("uuid must be different than empty")
	}
	return trainService.repository.DeleteByID(context, uuid)
}

// UpdateTrain updates a train by their UUID
func (trainService *TrainService) UpdateTrain(context context.Context, uuid string, updateTrain *models.UpdateTrain) (*models.Train, error) {
	log.Println("DB lookup for UUID:", uuid)
	train, err := trainService.repository.GetByID(context, uuid)
	log.Printf("FOUND TRAIN: %+v\n", train)
	if err != nil {
		return nil, fmt.Errorf("get train by id: %w", err)
	}
	switch updateTrain.Type {
	case models.TrainTypeRegional, models.TrainTypeIntercity, models.TrainTypeHighSpeed:
		train.Type = updateTrain.Type
	default:
		return nil, fmt.Errorf("unknown train type: %s", updateTrain.Type)
	}

	switch updateTrain.Status {
	case models.StatusActive, models.StatusInactive:
		train.Status = updateTrain.Status
	default:
		return nil, fmt.Errorf("unknown train status: %s", updateTrain.Status)
	}

	if updateTrain.Number != "" {
		train.Number = updateTrain.Number
	}

	if updateTrain.Capacity > 0 {
		train.Capacity = updateTrain.Capacity
	}

	errorUpdating := trainService.repository.Update(context, train)
	if errorUpdating != nil {
		return nil, fmt.Errorf("update train: %w", errorUpdating)
	}

	return train, nil
}

// PublishArrival a function for publishing a message for status of the train
func (trainService *TrainService) PublishArrival(trainUUID string, scheduleID int64) error {
	event := events.TrainEvent{
		Event: "arrived",
		Time:  time.Now().Format(time.RFC3339),
	}
	payload, _ := json.Marshal(event)
	topic := topics.TrainEventsTopicFor(trainUUID, scheduleID)
	token := trainService.mqttClient.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}
