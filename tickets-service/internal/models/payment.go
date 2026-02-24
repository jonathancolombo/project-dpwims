package models

// Payment represents a payment entity with its attributes.
type Payment struct {
	UUID              string        `json:"uuid"`
	TicketID          string        `json:"ticket_id"`
	Amount            int64         `json:"amount"`
	PaymentMethod     PaymentMethod `json:"payment_method"`
	ProviderReference string        `json:"provider_reference"`
}

type PaymentMethod string

const (
	PaymentMethodCreditCard   PaymentMethod = "credit_card"
	PaymentMethodPayPal       PaymentMethod = "banknotes"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
)
