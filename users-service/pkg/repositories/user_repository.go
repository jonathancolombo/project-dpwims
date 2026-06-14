package repositories

import (
	"context"
	"users-service/pkg/models"
)

// IUserRepository defines the interface for a user repositories.
type IUserRepository interface {
	Create(context context.Context, user *models.User) (*models.User, error)
	DeleteByID(context context.Context, id int64) error
	GetByID(context context.Context, id int64) (*models.User, error)
	GetAll(context context.Context) ([]*models.User, error)
	Update(context context.Context, user *models.User) error
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}
