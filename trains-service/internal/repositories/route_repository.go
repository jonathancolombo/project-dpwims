package repositories

import (
	"context"
	"trains-service/internal/models"
)

// IRouteRepository defines the interface for managing Route entities in the data source.
type IRouteRepository interface {
	Create(context context.Context, route *models.Route) (*models.Route, error)
	DeleteByID(context context.Context, id int64) error
	GetByID(context context.Context, id int64) (*models.Route, error)
	GetAll(context context.Context) ([]*models.Route, error)
	Update(context context.Context, route *models.Route) error
}
