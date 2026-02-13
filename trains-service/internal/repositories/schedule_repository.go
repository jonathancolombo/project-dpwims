package repositories

import (
	"context"
	"trains-service/internal/models"
)

// IScheduleRepository defines the interface for managing Schedule entities in the data source.
type IScheduleRepository interface {
	Create(context context.Context, route *models.Schedule) (*models.Schedule, error)
	DeleteByID(context context.Context, id int64) error
	GetByID(context context.Context, id int64) (*models.Schedule, error)
	GetAll(context context.Context) ([]*models.Schedule, error)
	Update(context context.Context, route *models.Schedule) error
}
