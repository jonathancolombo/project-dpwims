package repositories

import (
	"context"
	"tickets-service/internal/models"
)

// ITicketRepository defines the interface for managing Ticket entities in the data source.
type ITicketRepository interface {
	Create(context context.Context, ticket *models.Ticket) (*models.Ticket, error)
	DeleteByID(context context.Context, uuid string) error
	GetByID(context context.Context, uuid string) (*models.Ticket, error)
	GetAll(context context.Context) ([]*models.Ticket, error)
	Update(context context.Context, ticket *models.Ticket) error
}
