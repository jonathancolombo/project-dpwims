package models

// Payment represents a payment entity with its attributes.
type Payment struct {
	UUID              string        `json:"uuid"`
	TicketID          string        `json:"ticket_id:omitempty"`
	Amount            int64         `json:"amount:omitempty"`
	PaymentMethod     PaymentMethod `json:"payment_method:omitempty"`
	ProviderReference string        `json:"provider_reference"`
}

// PaymentMethod defines the type for different payment methods.
type PaymentMethod string

const (
	PaymentMethodCreditCard   PaymentMethod = "credit_card"
	PaymentMethodPayPal       PaymentMethod = "banknotes"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
)

// UpdatePayment represents the fields that can be updated for a payment entity.
type UpdatePayment struct {
	TicketID          string        `json:"ticket_id"`
	Amount            int64         `json:"amount"`
	PaymentMethod     PaymentMethod `json:"payment_method"`
	ProviderReference string        `json:"provider_reference"`
}
