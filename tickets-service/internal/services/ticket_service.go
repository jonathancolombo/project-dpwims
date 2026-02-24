package services

import (
	"context"
	"errors"
	"fmt"
	"tickets-service/internal/models"
	"tickets-service/internal/repositories"

	"github.com/google/uuid"
)

// TicketService defines the interface for managing Ticket entities.
type TicketService struct {
	repository repositories.ITicketRepository
}

// NewTicketService creates a new TicketService instance
func NewTicketService(repository repositories.ITicketRepository) *TicketService {
	return &TicketService{
		repository: repository,
	}
}

// CreateTicket creates a new ticket
func (ticketService *TicketService) CreateTicket(context context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	if ticket == nil {
		return nil, errors.New("ticket is nil")
	}

	if ticket.TrainUUID == "" {
		return nil, errors.New("train id is empty")
	}

	if ticket.ScheduleID <= 0 {
		return nil, errors.New("schedule id cannot be minor or equal to zero")
	}

	if ticket.SeatNumber <= 0 {
		return nil, errors.New("seat number cannot be minor or equal to zero")
	}

	if ticket.Price <= 0 {
		return nil, errors.New("price cannot be minor or equal to zero")
	}

	if ticket.Status == "" {
		return nil, errors.New("status is empty")
	}

	ticket.UUID = uuid.NewString()
	return ticketService.repository.Create(context, ticket)
}

// GetTicket retrieves a ticket by their UUID
func (ticketService *TicketService) GetTicket(context context.Context, uuid string) (*models.Ticket, error) {
	if uuid == "" {
		return nil, errors.New("uuid must be different than empty")
	}
	return ticketService.repository.GetByID(context, uuid)
}

// GetAllTickets retrieves all tickets
func (ticketService *TicketService) GetAllTickets(context context.Context) ([]*models.Ticket, error) {
	if ticketService.repository == nil {
		return nil, errors.New("repository is nil")
	}

	return ticketService.repository.GetAll(context)
}

// DeleteTicketByID deletes a ticket by their UUID
func (ticketService *TicketService) DeleteTicketByID(context context.Context, uuid string) error {
	if uuid == "" {
		return errors.New("uuid must be different than empty")
	}
	return ticketService.repository.DeleteByID(context, uuid)
}

// UpdateTicket updates a ticket
func (ticketService *TicketService) UpdateTicket(context context.Context, uuid string, updatedTicket *models.UpdateTicket) (*models.Ticket, error) {
	if uuid == "" {
		return nil, errors.New("uuid must be different than empty")
	}

	if updatedTicket == nil {
		return nil, errors.New("ticket is nil")
	}

	ticket, err := ticketService.repository.GetByID(context, uuid)
	if err != nil {
		return nil, fmt.Errorf("get ticket by id: %w", err)
	}

	if updatedTicket.TrainUUID != "" {
		ticket.TrainUUID = updatedTicket.TrainUUID
	}

	if updatedTicket.ScheduleID > 0 {
		ticket.ScheduleID = updatedTicket.ScheduleID
	}

	if updatedTicket.SeatNumber > 0 {
		ticket.SeatNumber = updatedTicket.SeatNumber
	}

	if updatedTicket.Price > 0 {
		ticket.Price = updatedTicket.Price
	}

	switch updatedTicket.Status {
	case models.TicketStatusBooked, models.TicketStatusCanceled, models.TicketStatusUsed:
		ticket.Status = updatedTicket.Status
	default:
		return nil, fmt.Errorf("unknown ticket status: %s", updatedTicket.Status)
	}

	errorUpdating := ticketService.repository.Update(context, ticket)
	if errorUpdating != nil {
		return nil, fmt.Errorf("update ticket: %w", errorUpdating)
	}

	return ticket, nil
}
