package services

import (
	"context"
	"errors"
	"fmt"
	"trains-service/internal/models"
	"trains-service/internal/repositories"
)

// StationService defines the interface for managing Station entities.
type StationService struct {
	repository repositories.IStationRepository
}

// NewStationService creates a new StationService instance
func NewStationService(repository repositories.IStationRepository) *StationService {
	return &StationService{
		repository: repository,
	}
}

// CreateStation creates a new station
func (stationService *StationService) CreateStation(context context.Context, station *models.Station) (*models.Station, error) {

	if station == nil {
		return nil, errors.New("station is nil")
	}

	if station.Name == "" {
		return nil, errors.New("station name is empty")
	}

	if station.City == "" {
		return nil, errors.New("station city is empty")
	}

	if station.Region == "" {
		return nil, errors.New("station region is empty")
	}

	if station.Status == "" {
		fmt.Println("station status is empty, setting default status to active")
		station.Status = models.StatusActive
	}

	return stationService.repository.Create(context, station)
}

// GetStation retrieves a station by their id
func (stationService *StationService) GetStation(context context.Context, id int64) (*models.Station, error) {
	if id <= 0 {
		return nil, errors.New("id must be greater than 0")
	}
	return stationService.repository.GetByID(context, id)
}

// GetAllStations retrieves all stations
func (stationService *StationService) GetAllStations(context context.Context) ([]*models.Station, error) {
	if stationService.repository == nil {
		return nil, errors.New("repositories is nil")
	}

	return stationService.repository.GetAll(context)
}

// DeleteStation deletes a station by their id
func (stationService *StationService) DeleteStation(context context.Context, id int64) error {
	if id <= 0 {
		return errors.New("id must be greater than 0")
	}

	return stationService.repository.DeleteByID(context, id)
}

// UpdateStation updates a station
func (stationService *StationService) UpdateStation(
	ctx context.Context,
	id int64,
	update *models.UpdateStation,
) (*models.Station, error) {

	station, err := stationService.repository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get station by id: %w", err)
	}

	// Validazione status
	switch update.Status {
	case models.StatusActive, models.StatusInactive:
		station.Status = update.Status
	default:
		return nil, fmt.Errorf("unknown station status: %s", update.Status)
	}

	// Aggiornamento campi
	station.Name = update.Name
	station.City = update.City
	station.Region = update.Region

	// Salvataggio
	if err := stationService.repository.Update(ctx, station); err != nil {
		return nil, fmt.Errorf("update station: %w", err)
	}

	return station, nil
}
