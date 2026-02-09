package services

import (
	"context"
	"trains-service/internal/models"
	"trains-service/internal/repositories"
)

// TrainService defines the interface for managing Train entities.
type TrainService struct {
	repository repositories.ITrainRepository
}

// NewTrainService creates a new TrainService instance
func NewTrainService(repository repositories.ITrainRepository) *TrainService {
	return &TrainService{
		repository: repository,
	}
}

// CreateTrain creates a new train
func (trainService *TrainService) CreateTrain(context context.Context, train *models.Train) (*models.Train, error) {

	return trainService.repository.Create(context, train)
}

// GetTrain retrieves a train by their ID
func (trainService *TrainService) GetTrain(context context.Context, id int64) (*models.Train, error) {
	return trainService.repository.GetByID(context, id)
}

// GetAllTrains retrieves all trains
func (trainService *TrainService) GetAllTrains(context context.Context) ([]*models.Train, error) {
	return trainService.repository.GetAll(context)
}

// DeleteTrainByID deletes a train by their ID
func (trainService *TrainService) DeleteTrainByID(context context.Context, id int64) error {
	return trainService.repository.DeleteByID(context, id)
}

// UpdateTrain updates a train by their ID
func (trainService *TrainService) UpdateTrain(context context.Context, id int64, train *models.Train) (*models.Train, error) {
	return train, nil
}
