package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"tickets-service/internal/models"
	"tickets-service/internal/repositories"
	"tickets-service/internal/services"

	"github.com/go-chi/chi/v5"
)

const KeyContentType = "Content-Type"
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
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

// GetTicket a handlers method to get a ticket by id from repositories memory
func (ticketHandler *TicketHandler) GetTicket(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "uuid")

	ticket, err := ticketHandler.service.GetTicket(request.Context(), idStr)

	if err != nil || ticket == nil {
		http.Error(writer, errorMessageTicketNotFound, http.StatusNotFound)
		return
	}

	requesterID := request.Context().Value("userID").(int64)
	role := request.Context().Value("role").(string)

	if role != "admin" && ticket.UserId != requesterID {
		http.Error(writer, "forbidden", http.StatusForbidden)
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

	ticket, err := ticketHandler.service.GetTicket(request.Context(), idString)
	if err != nil || ticket == nil {
		http.Error(writer, errorMessageTicketNotFound, http.StatusNotFound)
		return
	}

	requesterID := request.Context().Value("userID").(int64)
	role := request.Context().Value("role").(string)

	if role != "admin" && ticket.UserId != requesterID {
		http.Error(writer, "forbidden", http.StatusForbidden)
		return
	}

	err = ticketHandler.service.DeleteTicketByID(request.Context(), idString)
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

func (ticketHandler *TicketHandler) GetTicketsByUserID(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")

	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(writer, "invalid user id", http.StatusBadRequest)
		return
	}

	requesterID := request.Context().Value("userID").(int64)
	role := request.Context().Value("role").(string)

	if role != "admin" && userID != requesterID {
		http.Error(writer, "forbidden", http.StatusForbidden)
		return
	}

	tickets, err := ticketHandler.service.GetTicketsByUserID(request.Context(), userID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(tickets)
	if err != nil {
		return
	}
}
