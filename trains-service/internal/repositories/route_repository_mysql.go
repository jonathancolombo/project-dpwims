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

// DeleteByID is a method to delete the right route using id field
func (mySqlRouteRepository *MySQLRouteRepository) DeleteByID(context context.Context, id int64) error {
	query := `DELETE FROM routes WHERE id = ?`

	result, err := mySqlRouteRepository.database.ExecContext(context, query, id)
	if err != nil {
		return fmt.Errorf("delete route: %w", err)
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

// GetByID is a method to find the right route using id field
func (mySqlRouteRepository *MySQLRouteRepository) GetByID(context context.Context, id int64) (*models.Route, error) {
	query := `
        SELECT id, train_id, departure_station, arrival_station, distance
        FROM routes
        WHERE id = ?
    `
	row := mySqlRouteRepository.database.QueryRowContext(context, query, id)
	var route models.Route
	err := row.Scan(&route.ID, &route.TrainId, &route.DepartureStation, &route.ArrivalStation, &route.Distance)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrRouteNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan route: %w", err)
	}

	return &route, nil
}

// GetAll retrieves all routes into a slice
func (mySqlRouteRepository *MySQLRouteRepository) GetAll(context context.Context) ([]*models.Route, error) {
	query := `
        SELECT id, train_id, departure_station, arrival_station, distance
        FROM routes
    `

	rows, err := mySqlRouteRepository.database.QueryContext(context, query)
	if err != nil {
		return nil, fmt.Errorf("query all routes: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var routes []*models.Route
	for rows.Next() {
		var route models.Route
		if err := rows.Scan(&route.ID, &route.TrainId, &route.DepartureStation, &route.ArrivalStation, &route.Distance); err != nil {
			return nil, fmt.Errorf("scan route: %w", err)
		}
		routes = append(routes, &route)
	}

	return routes, nil
}

// Update a method to update a route by id from repository memory
func (mySqlRouteRepository *MySQLRouteRepository) Update(context context.Context, route *models.Route) error {
	query := `
        UPDATE routes
        SET train_id = ?, departure_station = ?, arrival_station = ?, distance = ?
        WHERE id = ?
    `
	_, err := mySqlRouteRepository.database.ExecContext(context, query,
		route.TrainId,
		route.DepartureStation,
		route.ArrivalStation,
		route.Distance,
	)
	if err != nil {
		return fmt.Errorf("route error: %w", err)
	}

	return nil
}
