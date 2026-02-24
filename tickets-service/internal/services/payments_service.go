package services

import (
	"context"
	"errors"
	"fmt"
	"tickets-service/internal/models"
	"tickets-service/internal/repositories"

	"github.com/google/uuid"
)

// PaymentService defines the interface for managing Payment entities.
type PaymentService struct {
	repository repositories.IPaymentRepository
}

// NewPaymentService creates a new PaymentService instance
func NewPaymentService(repository repositories.IPaymentRepository) *PaymentService {
	return &PaymentService{
		repository: repository,
	}
}

// CreatePayment creates a new payment
func (paymentService *PaymentService) CreatePayment(context context.Context, payment *models.Payment) (*models.Payment, error) {
	if payment == nil {
		return nil, errors.New("payment is nil")
	}
	if payment.TicketID == "" {
		return nil, errors.New("ticket id is empty")
	}
	if payment.Amount <= 0 {
		return nil, errors.New("amount cannot be minor or equal to zero")
	}
	if payment.PaymentMethod == "" {
		return nil, errors.New("payment method is empty")
	}
	if payment.ProviderReference == "" {
		return nil, errors.New("provider reference is empty")
	}
	payment.UUID = uuid.NewString()
	return paymentService.repository.Create(context, payment)
}

// GetPayment retrieves a payment by their UUID
func (paymentService *PaymentService) GetPayment(context context.Context, uuid string) (*models.Payment, error) {
	if uuid == "" {
		return nil, errors.New("uuid must be different than empty")
	}
	return paymentService.repository.GetByID(context, uuid)
}

// GetAllPayments retrieves all tickets
func (paymentService *PaymentService) GetAllPayments(context context.Context) ([]*models.Payment, error) {
	if paymentService.repository == nil {
		return nil, errors.New("repository is nil")
	}

	return paymentService.repository.GetAll(context)
}

// DeletePaymentByID deletes a payment by their UUID
func (paymentService *PaymentService) DeletePaymentByID(context context.Context, uuid string) error {
	if uuid == "" {
		return errors.New("uuid must be different than empty")
	}
	return paymentService.repository.DeleteByID(context, uuid)
}

// UpdatePayment updates a payment by their UUID
func (paymentService *PaymentService) UpdatePayment(context context.Context, uuid string, updatePayment *models.UpdatePayment) (*models.Payment, error) {
	if uuid == "" {
		return nil, errors.New("uuid must be different than empty")
	}

	if updatePayment == nil {
		return nil, errors.New("payment is nil")
	}

	payment, err := paymentService.repository.GetByID(context, uuid)
	if err != nil {
		return nil, fmt.Errorf("get payment by id: %w", err)
	}

	if updatePayment.TicketID != "" {
		payment.TicketID = updatePayment.TicketID
	}

	if updatePayment.Amount > 0 {
		payment.Amount = updatePayment.Amount
	}

	if updatePayment.PaymentMethod != "" {
		payment.PaymentMethod = updatePayment.PaymentMethod
	}

	if updatePayment.ProviderReference != "" {
		payment.ProviderReference = updatePayment.ProviderReference
	}

	errorUpdating := paymentService.repository.Update(context, payment)
	if errorUpdating != nil {
		return nil, fmt.Errorf("update payment: %w", errorUpdating)
	}

	return payment, nil
}
