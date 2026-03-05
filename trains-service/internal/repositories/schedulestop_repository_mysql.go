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

// NewMySQLStopScheduleRepository NewMySqlScheduleStopRepository initializes a new MySqlScheduleStopRepository with the provided database connection.
func NewMySQLStopScheduleRepository(db *sql.DB) *MySqlScheduleStopRepository {
	return &MySqlScheduleStopRepository{database: db}
}

// Create a method to create a stop schedule and save into a db
func (mySqlScheduleStopRepository *MySqlScheduleStopRepository) Create(ctx context.Context, stop *models.ScheduleStop) (*models.ScheduleStop, error) {
	if stop == nil {
		return nil, errors.New("schedule stop is nil")
	}

	query := `
        INSERT INTO schedules_stops 
        (schedule_id, station_id, stop_order, arrival_time, departure_time)
        VALUES (?, ?, ?, ?, ?)
    `

	stmt, err := mySqlScheduleStopRepository.database.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("prepare statement: %w", err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	result, err := stmt.Exec(
		stop.ScheduleID,
		stop.StationID,
		stop.StopOrder,
		stop.ArrivalTime,
		stop.DepartureTime,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert stop schedule: %w", err)
	}

	id, err := result.LastInsertId()
	if err == nil {
		stop.ID = id
	}

	return stop, nil
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

func (mySqlScheduleStopRepository *MySqlScheduleStopRepository) GetStopsBySchedule(ctx context.Context, scheduleId int64) ([]*models.ScheduleStop, error) {
	query := `
        SELECT 
            ss.id,
            ss.schedule_id,
            ss.station_id,
            s.name AS station_name,
            ss.stop_order,
            ss.arrival_time,
            ss.departure_time
        FROM schedules_stops ss
        JOIN stations s ON ss.station_id = s.id
        WHERE ss.schedule_id = ?
        ORDER BY ss.stop_order ASC
    `

	rows, err := mySqlScheduleStopRepository.database.QueryContext(ctx, query, scheduleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stops []*models.ScheduleStop

	for rows.Next() {
		var stop models.ScheduleStop
		if err := rows.Scan(
			&stop.ID,
			&stop.ScheduleID,
			&stop.StationID,
			&stop.StationName,
			&stop.StopOrder,
			&stop.ArrivalTime,
			&stop.DepartureTime,
		); err != nil {
			return nil, err
		}

		stops = append(stops, &stop)
	}

	return stops, nil
}
