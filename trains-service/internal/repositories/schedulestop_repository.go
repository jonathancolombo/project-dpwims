package repositories

import (
	"context"
	"trains-service/internal/models"
)

// IScheduleStopRepository defines the interface for managing ScheduleStop entities in the data source.
type IScheduleStopRepository interface {
	Create(context context.Context, scheduleStop *models.ScheduleStop) (*models.ScheduleStop, error)
	DeleteByID(context context.Context, id int64) error
	GetByID(context context.Context, id int64) (*models.ScheduleStop, error)
	GetAll(context context.Context) ([]*models.ScheduleStop, error)
	Update(context context.Context, scheduleStop *models.ScheduleStop) error
}
