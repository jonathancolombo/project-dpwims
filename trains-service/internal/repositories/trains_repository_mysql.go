package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"trains-service/internal/models"
)

// MySQLTrainRepository provides methods for CRUD operations on the 'trains' table in a MySQL database.
type MySQLTrainRepository struct {
	database *sql.DB
}

// NewMySQLRepositoryTrains initializes a new MySQLTrainRepository with the provided database connection.
func NewMySQLRepositoryTrains(db *sql.DB) *MySQLTrainRepository {
	return &MySQLTrainRepository{database: db}
}

// Create a method to create a train and save into a db
func (mySqlTrainRepository *MySQLTrainRepository) Create(context context.Context, train *models.Train) (*models.Train, error) {
	if train == nil {
		return nil, errors.New("train is nil")
	}

	query := `INSERT INTO trains (train_number, type, capacity, status) VALUES (?, ?, ?, ?)`

	statement, err := mySqlTrainRepository.database.PrepareContext(context, query)

	if err != nil {
		_ = fmt.Errorf("prepare statement: %w", err)
		return nil, err
	}

	defer statement.Close()

	result, err := statement.Exec(train.Number, strings.ToLower(train.Type), train.Capacity, strings.ToLower(train.Status))
	if err != nil {
		return nil, fmt.Errorf("failed to insert train: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to insert id: %w", err)
	}
	train.ID = id
	return train, nil
}

// GetByID is a method to find the right train using id field
func (mySqlTrainRepository *MySQLTrainRepository) GetByID(context context.Context, id int64) (*models.Train, error) {
	if id <= 0 {
		return nil, errors.New("id must be greater than 0")
	}
	query := `SELECT train_number, type, capacity, status FROM trains WHERE id = ?`
	rows, err := mySqlTrainRepository.database.QueryContext(context, query, id)
	if err != nil {
		return nil, fmt.Errorf("query train by id: %w", err)
	}

	defer rows.Close()
	var train models.Train
	errorScan := rows.Scan(&train.Number, &train.Type, &train.Capacity, &train.Status)
	if errors.Is(errorScan, sql.ErrNoRows) {
		return nil, fmt.Errorf("scan train: %w", errorScan)
	}

	if errorScan != nil {
		return nil, fmt.Errorf("scan train: %w", errorScan)
	}
	return &train, nil
}

// GetAll retrieves all trains into a slice
func (mySqlTrainRepository *MySQLTrainRepository) GetAll(context context.Context) ([]*models.Train, error) {
	query := `
        SELECT id, train_number, type, capacity, status
        FROM trains
    `

	rows, err := mySqlTrainRepository.database.QueryContext(context, query)
	if err != nil {
		return nil, fmt.Errorf("query all trains: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var trains []*models.Train
	for rows.Next() {
		var train models.Train
		if err := rows.Scan(&train.ID, &train.Number, &train.Type, &train.Capacity, &train.Status); err != nil {
			return nil, fmt.Errorf("scan train: %w", err)
		}
		trains = append(trains, &train)
	}

	return trains, nil
}

// DeleteByID delete a train by his id
func (mySqlTrainRepository *MySQLTrainRepository) DeleteByID(context context.Context, id int64) error {
	if id <= 0 {
		return errors.New("id must be greater than 0")
	}

	query := `DELETE FROM trains WHERE id = ?`

	result, err := mySqlTrainRepository.database.ExecContext(context, query, id)
	if err != nil {
		return fmt.Errorf("delete train: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rows == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

// Update a train by his id
func (mySqlTrainRepository *MySQLTrainRepository) Update(context context.Context, train *models.Train) error {
	query := `
        UPDATE trains
        SET train_number = ?, type = ?, capacity = ?, status = ? 
        WHERE id = ?
    `
	_, err := mySqlTrainRepository.database.ExecContext(context, query,
		train.Number,
		train.Type,
		train.Capacity,
		train.Status,
		train.ID,
	)

	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	return nil
}
