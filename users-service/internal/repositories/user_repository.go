package repositories

import (
	"context"
	"user-service/internal/models"
)

// IUserRepository defines the interface for a user repository.
// All methods accept a context because the underlying implementation
// may involve I/O operations (SQL, network, etc.).
type IUserRepository interface {
	Create(context context.Context, user *models.User) (*models.User, error)
	DeleteByID(context context.Context, id int64) error
	GetByID(context context.Context, id int64) (*models.User, error)
	GetAll(context context.Context) ([]*models.User, error)
	Update(context context.Context, user *models.User) error
}
