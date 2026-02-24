package repository

import "context"

type ITicketRepository interface {
	Create(context context.Context, ticket *models.Ticket) (*models.Ticket, error)
	DeleteByID(context context.Context, id string) error
}
