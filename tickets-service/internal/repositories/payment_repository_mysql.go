package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"tickets-service/internal/models"
)

// MySQLPaymentRepository provides methods for CRUD operations on the 'payments' table in a MySQL database.
type MySQLPaymentRepository struct {
	database *sql.DB
}

// NewMySQLRepositoryPayments initializes a new MySQLPaymentRepository with the provided database connection.
func NewMySQLRepositoryPayments(db *sql.DB) *MySQLPaymentRepository {
	return &MySQLPaymentRepository{database: db}
}

// Create a method to create a payment and save into a db
func (mySqlPaymentRepository *MySQLPaymentRepository) Create(context context.Context, payment *models.Payment) (*models.Payment, error) {
	if payment == nil {
		return nil, errors.New("payment is nil")
	}

	query := `INSERT INTO payments (ticket_id, amount, payment_method, provider_reference) VALUES (?, ?, ?, ?)`
	statement, err := mySqlPaymentRepository.database.PrepareContext(context, query)

	if err != nil {
		_ = fmt.Errorf("prepare statement: %w", err)
		return nil, err
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			return
		}
	}(statement)

	_, err = statement.Exec(payment.TicketID, payment.Amount, payment.PaymentMethod, payment.ProviderReference)
	if err != nil {
		return nil, fmt.Errorf("failed to insert payment: %w", err)
	}
	return payment, nil
}

// GetByID is a method to find the right payment using id field
func (mySqlPaymentRepository *MySQLPaymentRepository) GetByID(context context.Context, uuid string) (*models.Payment, error) {
	if uuid == "" {
		return nil, errors.New("uuid must be greater than 0")
	}
	query := `SELECT ticket_id, amount, payment_method, provider_reference FROM payments WHERE uuid = ?`
	rows := mySqlPaymentRepository.database.QueryRowContext(context, query, uuid)

	var payment models.Payment
	errorScan := rows.Scan(&payment.TicketID, &payment.Amount, &payment.PaymentMethod, &payment.ProviderReference)

	if errors.Is(errorScan, sql.ErrNoRows) {
		return nil, fmt.Errorf("payment with uuid %s not found: %w", uuid, errorScan)
	}

	if errorScan != nil {
		return nil, fmt.Errorf("failed to scan payment: %w", errorScan)
	}

	return &payment, nil
}

// GetAll retrieves all payments into a slice
func (mySqlPaymentRepository *MySQLPaymentRepository) GetAll(context context.Context) ([]*models.Payment, error) {
	query := `SELECT uuid, ticket_id, amount, payment_method, provider_reference FROM payments`
	rows, err := mySqlPaymentRepository.database.QueryContext(context, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query payments: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var payments []*models.Payment
	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(&payment.UUID, &payment.TicketID, &payment.Amount, &payment.PaymentMethod, &payment.ProviderReference)
		if err != nil {
			return nil, fmt.Errorf("failed to scan payment: %w", err)
		}
		payments = append(payments, &payment)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over payments: %w", err)
	}

	return payments, nil
}

// DeleteByID delete a payment by his id
func (mySqlPaymentRepository *MySQLPaymentRepository) DeleteByID(context context.Context, uuid string) error {
	if uuid == "" {
		return errors.New("uuid must be greater than 0")
	}
	query := `DELETE FROM payments WHERE uuid = ?`
	statement, err := mySqlPaymentRepository.database.PrepareContext(context, query)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			return
		}
	}(statement)

	result, err := statement.ExecContext(context, uuid)
	if err != nil {
		return fmt.Errorf("failed to delete payment: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no payment found with uuid %s", uuid)
	}

	return nil
}

// Update a payment by his id
func (mySqlPaymentRepository *MySQLPaymentRepository) Update(context context.Context, payment *models.Payment) error {
	if payment == nil {
		return errors.New("payment is nil")
	}

	query := `UPDATE payments SET ticket_id = ?, amount = ?, payment_method = ?, provider_reference = ? WHERE uuid = ?`
	result, err := mySqlPaymentRepository.database.ExecContext(context, query, payment.TicketID, payment.Amount, payment.PaymentMethod, payment.ProviderReference)
	if err != nil {
		return fmt.Errorf("failed to update ticket: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no payment found with uuid %s", payment.UUID)
	}

	return nil
}
