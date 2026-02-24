package repositories

import (
	"context"
	"tickets-service/internal/models"
)

// IPaymentRepository defines the interface for managing Payment entities in the data source.
type IPaymentRepository interface {
	Create(context context.Context, ticket *models.Payment) (*models.Payment, error)
	DeleteByID(context context.Context, uuid string) error
	GetByID(context context.Context, uuid string) (*models.Payment, error)
	GetAll(context context.Context) ([]*models.Payment, error)
	Update(context context.Context, ticket *models.Payment) error
}
