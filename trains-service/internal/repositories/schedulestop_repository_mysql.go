package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"trains-service/internal/models"
)

// MySqlScheduleStopRepository is a struct that implements the IScheduleStopRepository interface using a MySQL database as the data source.
type MySqlScheduleStopRepository struct {
	database *sql.DB
}

// NewMySqlScheduleStopRepository initializes a new MySqlScheduleStopRepository with the provided database connection.
func NewMySQLStopScheduleRepository(db *sql.DB) *MySqlScheduleStopRepository {
	return &MySqlScheduleStopRepository{database: db}
}

// Create a method to create a stop schedule and save into a db
func (mySqlScheduleStopRepository *MySqlScheduleStopRepository) Create(context context.Context, schedule *models.ScheduleStop) (*models.ScheduleStop, error) {
	if schedule == nil {
		return nil, errors.New("schedule is nil")
	}
	query := `INSERT INTO schedules_stops 
    (id, schedule_id, station_id, station_name, stop_order, arrival_time, departure_time) VALUES (?, ?, ?, ?, ?, ?, ?)`

	statement, err := mySqlScheduleStopRepository.database.PrepareContext(context, query)

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

	_, err = statement.Exec(schedule.ID, schedule.ScheduleID, schedule.StationID, schedule.StationName, schedule.StopOrder, schedule.ArrivalTime, schedule.DepartureTime)
	if err != nil {
		return nil, fmt.Errorf("failed to insert stop schedule: %w", err)
	}

	return schedule, nil
}

// DeleteByID is a method to delete the right stop schedule using id field
func (mySqlScheduleStopRepository *MySqlScheduleStopRepository) DeleteByID(context context.Context, id int64) error {
	query := `DELETE FROM schedules_stops WHERE id = ?`

	result, err := mySqlScheduleStopRepository.database.ExecContext(context, query, id)
	if err != nil {
		return fmt.Errorf("delete stop schedule: %w", err)
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

// GetByID is a method to find the right stop schedule using id field
func (mySqlScheduleStopRepository *MySqlScheduleStopRepository) GetByID(context context.Context, id int64) (*models.ScheduleStop, error) {
	query := `
        SELECT id, schedule_id, station_id, station_name, stop_order, arrival_time, departure_time
        FROM schedules_stops
        WHERE id = ?
    `
	row := mySqlScheduleStopRepository.database.QueryRowContext(context, query, id)
	var schedule models.ScheduleStop
	err := row.Scan(&schedule.ID, &schedule.ScheduleID, &schedule.StationID, &schedule.StationName, &schedule.StopOrder, &schedule.ArrivalTime, &schedule.DepartureTime)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrRouteNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan stop schedule: %w", err)
	}

	return &schedule, nil
}

// GetAll retrieves all stop schedule into a slice
func (mySqlScheduleStopRepository *MySqlScheduleStopRepository) GetAll(context context.Context) ([]*models.ScheduleStop, error) {
	query := `
        SELECT id, schedule_id, station_id, station_name, stop_order, arrival_time, departure_time
        FROM schedules_stops
    `

	rows, err := mySqlScheduleStopRepository.database.QueryContext(context, query)
	if err != nil {
		return nil, fmt.Errorf("query all stop schedules: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var schedules []*models.ScheduleStop
	for rows.Next() {
		var schedule models.ScheduleStop
		if err := rows.Scan(&schedule.ID, &schedule.ScheduleID, &schedule.StationID, &schedule.StationName, &schedule.StopOrder, &schedule.ArrivalTime, &schedule.DepartureTime); err != nil {
			return nil, fmt.Errorf("scan schedule: %w", err)
		}
		schedules = append(schedules, &schedule)
	}

	return schedules, nil
}

// Update a method to update a stop schedule by id from repositories memory
func (mySqlScheduleStopRepository *MySqlScheduleStopRepository) Update(context context.Context, schedule *models.ScheduleStop) error {
	query := `
        UPDATE schedules_stops
        SET schedule_id = ?, station_id = ? , station_name = ?, stop_order  = ?, arrival_time = ?, departure_time = ?
        WHERE id = ?
    `
	_, err := mySqlScheduleStopRepository.database.ExecContext(context, query,
		schedule.ScheduleID,
		schedule.StationID,
		schedule.StationName,
		schedule.StopOrder,
		schedule.ArrivalTime,
		schedule.DepartureTime,
	)
	if err != nil {
		return fmt.Errorf("stop schedule error: %w", err)
	}

	return nil
}
