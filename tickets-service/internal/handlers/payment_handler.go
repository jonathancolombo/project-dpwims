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

const errorMessagePaymentNotFound = "payment not found"

// PaymentHandler is responsible for handling HTTP requests related to Payment entities.
type PaymentHandler struct {
	service *services.PaymentService
}

// NewPaymentHandler to create an instance of PaymentHandler
func NewPaymentHandler(paymentService *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: paymentService}
}

// CreatePayment to manage http request to create a payment and save it into repositories memory
func (paymentHandler *PaymentHandler) CreatePayment(writer http.ResponseWriter, request *http.Request) {
	var payment models.Payment
	err := json.NewDecoder(request.Body).Decode(&payment)
	if err != nil {
		http.Error(writer, "invalid JSON body"+err.Error(), http.StatusBadRequest)
	}

	created, err := paymentHandler.service.CreatePayment(request.Context(), &payment)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(created)
}

// GetPayment a handlers method to get a payment by id from repository memory
func (paymentHandler *PaymentHandler) GetPayment(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "uuid")

	payment, err := paymentHandler.service.GetPayment(request.Context(), idStr)

	if err != nil || payment == nil {
		http.Error(writer, errorMessagePaymentNotFound, http.StatusNotFound)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(payment)
}

// GetAllPayments a handlers method to get all payments into repositories memory
func (paymentHandler *PaymentHandler) GetAllPayments(writer http.ResponseWriter, request *http.Request) {
	tickets, err := paymentHandler.service.GetAllPayments(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(tickets)
}

// DeletePayment a handlers method to delete a ticket by id from repositories memory
func (paymentHandler *PaymentHandler) DeletePayment(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "uuid")
	if idString == "" {
		http.Error(writer, errorMessageInvalidID, http.StatusBadRequest)
		return
	}

	err := paymentHandler.service.DeletePaymentByID(request.Context(), idString)
	if err != nil {
		http.Error(writer, errorMessagePaymentNotFound, http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

// UpdatePayment a handlers method to update a payment by id from repository memory
func (paymentHandler *PaymentHandler) UpdatePayment(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "uuid")

	if idString == "" {
		http.Error(writer, errorMessageInvalidID, http.StatusBadRequest)
		return
	}

	var updatePaymentRequest models.UpdatePayment
	if err := json.NewDecoder(request.Body).Decode(&updatePaymentRequest); err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}

	updatePayment, err := paymentHandler.service.UpdatePayment(request.Context(), idString, &updatePaymentRequest)
	if err != nil {
		if errors.Is(err, repositories.ErrTicketNotFound) {
			http.Error(writer, errorMessagePaymentNotFound, http.StatusNotFound)
			return
		}

		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(updatePayment)
	if err != nil {
		return
	}
}
