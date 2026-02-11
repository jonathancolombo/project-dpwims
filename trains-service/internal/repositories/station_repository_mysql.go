package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"trains-service/internal/models"
)

var ErrStationNotFound = errors.New("station not found")

// MySQLStationRepository provides methods for CRUD operations on the 'stations' table in a MySQL database.
type MySQLStationRepository struct {
	database *sql.DB
}

// NewMySQLRepositoryStation initializes a new MySQLStationRepository with the provided database connection.`
func NewMySQLRepositoryStation(db *sql.DB) *MySQLStationRepository {
	return &MySQLStationRepository{database: db}
}

// Create a method to create a station and save into a db
func (mySqlStationRepository *MySQLStationRepository) Create(context context.Context, station *models.Station) (*models.Station, error) {
	if station == nil {
		return nil, errors.New("station is nil")
	}

	query := `INSERT INTO stations (name, city, region, status) VALUES (?, ?, ?, ?)`

	statement, err := mySqlStationRepository.database.PrepareContext(context, query)

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

	_, err = statement.Exec(station.Name, station.City, station.Region, strings.ToLower(string(station.Status)))
	if err != nil {
		return nil, fmt.Errorf("failed to insert station: %w", err)
	}

	return station, nil
}

// DeleteByID is a method to delete the right station using id field
func (mySqlStationRepository *MySQLStationRepository) DeleteByID(context context.Context, id int64) error {
	query := `DELETE FROM stations WHERE id = ?`

	result, err := mySqlStationRepository.database.ExecContext(context, query, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rows == 0 {
		return ErrStationNotFound
	}

	return nil
}

// GetByID is a method to find the right station using id field
func (mySqlStationRepository *MySQLStationRepository) GetByID(context context.Context, id int64) (*models.Station, error) {
	query := `
        SELECT id, name, city, region, status
        FROM stations
        WHERE id = ?
    `
	row := mySqlStationRepository.database.QueryRowContext(context, query, id)
	var station models.Station
	err := row.Scan(&station.ID, &station.Name, &station.City, &station.Region, &station.Status)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrStationNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan station: %w", err)
	}

	return &station, nil
}

// GetAll retrieves all stations into a slice
func (mySqlStationRepository *MySQLStationRepository) GetAll(context context.Context) ([]*models.Station, error) {
	query := `
        SELECT id, name, city, region, status
        FROM stations
    `

	rows, err := mySqlStationRepository.database.QueryContext(context, query)
	if err != nil {
		return nil, fmt.Errorf("query all stations: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var stations []*models.Station
	for rows.Next() {
		var station models.Station
		if err := rows.Scan(&station.ID, &station.Name, &station.City, &station.Region, &station.Status); err != nil {
			return nil, fmt.Errorf("scan station: %w", err)
		}
		stations = append(stations, &station)
	}

	return stations, nil
}

// Update a method to update a station by id from repository memory
func (mySqlStationRepository *MySQLStationRepository) Update(context context.Context, station *models.Station) error {
	query := `
        UPDATE stations
        SET name = ?, city = ?, region = ?, status = ?
        WHERE id = ?
    `
	_, err := mySqlStationRepository.database.ExecContext(context, query,
		station.Name,
		station.City,
		station.Region,
		strings.ToLower(string(station.Status)),
	)
	if err != nil {
		return fmt.Errorf("station user: %w", err)
	}

	return nil
}
