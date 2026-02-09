package repositories

import (
	"context"
	"trains-service/internal/models"
)

// ITrainRepository defines the interface for managing Train entities in the data source.
type ITrainRepository interface {
	Create(context context.Context, train *models.Train) (*models.Train, error)
	DeleteByID(context context.Context, id int64) error
	GetByID(context context.Context, id int64) (*models.Train, error)
	GetAll(context context.Context) ([]*models.Train, error)
	Update(context context.Context, user *models.Train) error
}
