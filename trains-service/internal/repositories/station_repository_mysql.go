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

func (mySqlStationRepository *MySQLStationRepository) DeleteByID(context context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (mySqlStationRepository *MySQLStationRepository) GetByID(context context.Context, id int64) (*models.Station, error) {
	//TODO implement me
	panic("implement me")
}

func (mySqlStationRepository *MySQLStationRepository) GetAll(context context.Context) ([]*models.Station, error) {
	//TODO implement me
	panic("implement me")
}

func (mySqlStationRepository *MySQLStationRepository) Update(context context.Context, station *models.Station) (*models.Station, error) {
	//TODO implement me
	panic("implement me")
}

// NewMySQLRepositoryStation initializes a new MySQLStationRepository with the provided database connection.`
func NewMySQLRepositoryStation(db *sql.DB) *MySQLStationRepository {
	return &MySQLStationRepository{database: db}
}
