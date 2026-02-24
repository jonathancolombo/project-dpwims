package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"tickets-service/internal/models"
)

var ErrTicketNotFound = errors.New("ticket not found")

// MySQLTicketRepository provides methods for CRUD operations on the 'tickets' table in a MySQL database.
type MySQLTicketRepository struct {
	database *sql.DB
}

// NewMySQLRepositoryTickets initializes a new MySQLTicketRepository with the provided database connection.
func NewMySQLRepositoryTickets(db *sql.DB) *MySQLTicketRepository {
	return &MySQLTicketRepository{database: db}
}

// Create a method to create a ticket and save into a db
func (mySqlTicketRepository *MySQLTicketRepository) Create(context context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	if ticket == nil {
		return nil, errors.New("ticket is nil")
	}

	query := `INSERT INTO tickets (user_id, train_uuid, schedule_id, seat_number, price, status) VALUES (?, ?, ?, ?, ?, ?)`
	statement, err := mySqlTicketRepository.database.PrepareContext(context, query)

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

	_, err = statement.Exec(ticket.TrainUUID, ticket.ScheduleID, ticket.SeatNumber, ticket.Price, strings.ToLower(string(ticket.Status)))
	if err != nil {
		return nil, fmt.Errorf("failed to insert ticket: %w", err)
	}
	return ticket, nil
}

// GetByID is a method to find the right ticket using id field
func (mySqlTicketRepository *MySQLTicketRepository) GetByID(context context.Context, uuid string) (*models.Ticket, error) {
	if uuid == "" {
		return nil, errors.New("uuid must be greater than 0")
	}
	query := `SELECT user_id, train_uuid, schedule_id, seat_number, price, status FROM tickets WHERE uuid = ?`
	rows := mySqlTicketRepository.database.QueryRowContext(context, query, uuid)

	var ticket models.Ticket
	errorScan := rows.Scan(&ticket.TrainUUID, &ticket.ScheduleID, &ticket.SeatNumber, &ticket.Price, &ticket.Status)

	if errors.Is(errorScan, sql.ErrNoRows) {
		return nil, fmt.Errorf("ticket with uuid %s not found: %w", uuid, errorScan)
	}

	if errorScan != nil {
		return nil, fmt.Errorf("failed to scan ticket: %w", errorScan)
	}

	return &ticket, nil
}

// GetAll retrieves all tickets into a slice
func (mySqlTicketRepository *MySQLTicketRepository) GetAll(context context.Context) ([]*models.Ticket, error) {
	query := `SELECT user_id, train_uuid, schedule_id, seat_number, price, status FROM tickets`
	rows, err := mySqlTicketRepository.database.QueryContext(context, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query tickets: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var tickets []*models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		err := rows.Scan(&ticket.UserId, &ticket.TrainUUID, &ticket.ScheduleID, &ticket.SeatNumber, &ticket.Price, &ticket.Status)
		if err != nil {
			return nil, fmt.Errorf("failed to scan ticket: %w", err)
		}
		tickets = append(tickets, &ticket)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over tickets: %w", err)
	}

	return tickets, nil
}

// DeleteByID delete a ticket by his id
func (mySqlTicketRepository *MySQLTicketRepository) DeleteByID(context context.Context, uuid string) error {
	if uuid == "" {
		return errors.New("uuid must be greater than 0")
	}
	query := `DELETE FROM tickets WHERE uuid = ?`
	result, err := mySqlTicketRepository.database.ExecContext(context, query, uuid)
	if err != nil {
		return fmt.Errorf("failed to delete ticket: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no ticket found with uuid %s", uuid)
	}

	return nil
}

// Update a ticket by his id
func (mySqlTicketRepository *MySQLTicketRepository) Update(context context.Context, ticket *models.Ticket) error {
	if ticket == nil {
		return errors.New("ticket is nil")
	}

	query := `UPDATE tickets SET user_id = ?, train_uuid = ?, schedule_id = ?, seat_number = ?, price = ?, status = ? WHERE uuid = ?`
	result, err := mySqlTicketRepository.database.ExecContext(context, query, ticket.TrainUUID, ticket.ScheduleID, ticket.SeatNumber, ticket.Price, strings.ToLower(string(ticket.Status)), ticket.UUID)
	if err != nil {
		return fmt.Errorf("failed to update ticket: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no ticket found with uuid %s", ticket.UUID)
	}

	return nil
}
