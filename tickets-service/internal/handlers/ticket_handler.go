package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"tickets-service/internal/models"
	"tickets-service/internal/repositories"
	"tickets-service/internal/services"

	"github.com/go-chi/chi/v5"
)

const KeyContentType = "Content-TrainType"
const ValueAppJson = "application/json"
const errorMessageInvalidUUID = "invalid uuid"
const errorMessageTicketNotFound = "ticket not found"

// TicketHandler is responsible for handling HTTP requests related to Ticket entities.
type TicketHandler struct {
	service *services.TicketService
}

// NewTicketHandler to create an instance of TicketHandler
func NewTicketHandler(ticketService *services.TicketService) *TicketHandler {
	return &TicketHandler{service: ticketService}
}

// CreateTicket to manage api request to create a ticket
func (ticketHandler *TicketHandler) CreateTicket(writer http.ResponseWriter, request *http.Request) {
	var ticket models.Ticket
	err := json.NewDecoder(request.Body).Decode(&ticket)
	if err != nil {
		http.Error(writer, "invalid JSON body"+err.Error(), http.StatusBadRequest)
	}

	created, err := ticketHandler.service.CreateTicket(request.Context(), &ticket)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(created)
}

// GetTicket a handlers method to get a ticket by id from repositories memory
func (ticketHandler *TicketHandler) GetTicket(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "uuid")

	ticket, err := ticketHandler.service.GetTicket(request.Context(), idStr)

	if err != nil || ticket == nil {
		http.Error(writer, errorMessageTicketNotFound, http.StatusNotFound)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(ticket)
}

// GetAllTickets a handlers method to get all tickets into repositories memory
func (ticketHandler *TicketHandler) GetAllTickets(writer http.ResponseWriter, request *http.Request) {
	tickets, err := ticketHandler.service.GetAllTickets(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(tickets)
}

// DeleteTicket a handlers method to delete a ticket by id from repositories memory
func (ticketHandler *TicketHandler) DeleteTicket(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "uuid")
	if idString == "" {
		http.Error(writer, errorMessageInvalidUUID, http.StatusBadRequest)
		return
	}

	err := ticketHandler.service.DeleteTicketByID(request.Context(), idString)
	if err != nil {
		http.Error(writer, errorMessageTicketNotFound, http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

// UpdateTicket a handlers method to update a ticket by id from repositories memory
func (ticketHandler *TicketHandler) UpdateTicket(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "uuid")

	if idString == "" {
		http.Error(writer, errorMessageInvalidUUID, http.StatusBadRequest)
		return
	}

	var updateTicketRequest models.UpdateTicket
	if err := json.NewDecoder(request.Body).Decode(&updateTicketRequest); err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}

	updateTicket, err := ticketHandler.service.UpdateTicket(request.Context(), idString, &updateTicketRequest)
	if err != nil {
		if errors.Is(err, repositories.ErrTicketNotFound) {
			http.Error(writer, errorMessageTicketNotFound, http.StatusNotFound)
			return
		}

		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(updateTicket)
	if err != nil {
		return
	}
}
