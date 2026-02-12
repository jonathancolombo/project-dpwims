package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"trains-service/internal/models"
)

var ErrRouteNotFound = errors.New("station not found")

type MySQLRouteRepository struct {
	database *sql.DB
}

// NewMySQLRepositoryRoute initializes a new MySQLRouteRepository with the provided database connection.
func NewMySQLRepositoryRoute(db *sql.DB) *MySQLRouteRepository {
	return &MySQLRouteRepository{database: db}
}

// Create a method to create a route and save into a db
func (mySqlRouteRepository *MySQLRouteRepository) Create(context context.Context, route *models.Route) (*models.Route, error) {
	if route == nil {
		return nil, errors.New("route is nil")
	}
	query := `INSERT INTO routes 
    (train_id, departure_station, arrival_station, distance) VALUES (?, ?, ?, ?)`

	statement, err := mySqlRouteRepository.database.PrepareContext(context, query)

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

	_, err = statement.Exec(route.TrainId, route.DepartureStation, route.ArrivalStation, route.Distance)
	if err != nil {
		return nil, fmt.Errorf("failed to insert station: %w", err)
	}

	return route, nil
}
