package repositories

import (
	"context"
	"trains-service/internal/models"
)

// IStationRepository defines the interface for managing Station entities in the data source.
type IStationRepository interface {
	Create(context context.Context, station *models.Station) (*models.Station, error)
	DeleteByID(context context.Context, id int64) error
	GetByID(context context.Context, id int64) (*models.Station, error)
	GetAll(context context.Context) ([]*models.Station, error)
	Update(context context.Context, station *models.Station) (*models.Station, error)
}
