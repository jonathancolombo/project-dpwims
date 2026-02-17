package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"trains-service/internal/models"
)

var ErrScheduleNotFound = errors.New("schedule not found")

// MySQLScheduleRepository is a struct that implements the IScheduleRepository interface using a MySQL database as the data source.
type MySQLScheduleRepository struct {
	database *sql.DB
}

// NewMySQLScheduleRepository initializes a new MySQLScheduleRepository with the provided database connection.
func NewMySQLScheduleRepository(db *sql.DB) *MySQLScheduleRepository {
	return &MySQLScheduleRepository{database: db}
}

// Create a method to create a schedule and save into a db
func (mySqlScheduleRepository *MySQLScheduleRepository) Create(context context.Context, schedule *models.Schedule) (*models.Schedule, error) {
	if schedule == nil {
		return nil, errors.New("schedule is nil")
	}
	query := `INSERT INTO schedules 
    (id, train_id, station_id, departure_station, arrival_station, status, price) VALUES (?, ?, ?, ?, ?, ?, ?)`

	statement, err := mySqlScheduleRepository.database.PrepareContext(context, query)

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

	_, err = statement.Exec(schedule.ID, schedule.TrainID, schedule.StationID, schedule.Departure, schedule.Arrival, schedule.Status, schedule.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to insert schedule: %w", err)
	}

	return schedule, nil
}

// DeleteByID is a method to delete the right schedule using id field
func (mySqlScheduleRepository *MySQLScheduleRepository) DeleteByID(context context.Context, id int64) error {
	query := `DELETE FROM schedules WHERE id = ?`

	result, err := mySqlScheduleRepository.database.ExecContext(context, query, id)
	if err != nil {
		return fmt.Errorf("delete schedule: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rows == 0 {
		return ErrRouteNotFound
	}

	return nil
}

// GetByID is a method to find the right schedule using id field
func (mySqlScheduleRepository *MySQLScheduleRepository) GetByID(context context.Context, id int64) (*models.Schedule, error) {
	query := `
        SELECT id, train_id, station_id, departure_station, arrival_station, status, price
        FROM schedules
        WHERE id = ?
    `
	row := mySqlScheduleRepository.database.QueryRowContext(context, query, id)
	var schedule models.Schedule
	err := row.Scan(&schedule.ID, &schedule.TrainID, &schedule.StationID, &schedule.Departure, &schedule.Arrival, &schedule.Status, &schedule.Price)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrRouteNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan schedule: %w", err)
	}

	return &schedule, nil
}

// GetAll retrieves all schedule into a slice
func (mySqlScheduleRepository *MySQLScheduleRepository) GetAll(context context.Context) ([]*models.Schedule, error) {
	query := `
        SELECT id, train_id, station_id, departure_station, arrival_station, status, price
        FROM schedules
    `

	rows, err := mySqlScheduleRepository.database.QueryContext(context, query)
	if err != nil {
		return nil, fmt.Errorf("query all schedules: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var schedules []*models.Schedule
	for rows.Next() {
		var schedule models.Schedule
		if err := rows.Scan(&schedule.ID, &schedule.TrainID, &schedule.StationID, &schedule.Departure, &schedule.Arrival, &schedule.Status, &schedule.Price); err != nil {
			return nil, fmt.Errorf("scan schedule: %w", err)
		}
		schedules = append(schedules, &schedule)
	}

	return schedules, nil
}

// Update a method to update a schedule by id from repository memory
func (mySqlScheduleRepository *MySQLScheduleRepository) Update(context context.Context, schedule *models.Schedule) error {
	query := `
        UPDATE schedules
        SET train_id = ?, station_id = ? , departure_station = ?, arrival_station  = ?, status = ?, price = ?
        WHERE id = ?
    `
	_, err := mySqlScheduleRepository.database.ExecContext(context, query,
		schedule.TrainID,
		schedule.StationID,
		schedule.Departure,
		schedule.Arrival,
		schedule.Status,
		schedule.Price,
	)
	if err != nil {
		return fmt.Errorf("schedule error: %w", err)
	}

	return nil
}
