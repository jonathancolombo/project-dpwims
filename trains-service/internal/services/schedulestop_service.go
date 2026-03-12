package services

import (
	"context"
	"fmt"
	"trains-service/internal/models"
	"trains-service/internal/repositories"
)

// ScheduleStopService defines the interface for managing ScheduleStop entities.
type ScheduleStopService struct {
	repository repositories.IScheduleStopRepository
}

// NewStopScheduleService creates a new ScheduleStopService instance
func NewStopScheduleService(repository repositories.IScheduleStopRepository) *ScheduleStopService {
	return &ScheduleStopService{
		repository: repository,
	}
}

// CreateStopSchedule creates a new stop schedule
func (scheduleStopService *ScheduleStopService) CreateStopSchedule(context context.Context, scheduleStop *models.ScheduleStop) (*models.ScheduleStop, error) {
	stops, _ := scheduleStopService.repository.GetStopsBySchedule(context, scheduleStop.ScheduleID)

	scheduleStop.StopOrder = len(stops) + 1

	return scheduleStopService.repository.Create(context, scheduleStop)
}

// GetStopSchedule retrieves a stop schedule by their id
func (scheduleStopService *ScheduleStopService) GetStopSchedule(context context.Context, id int64) (*models.ScheduleStop, error) {
	if id <= 0 {
		return nil, nil
	}
	return scheduleStopService.repository.GetByID(context, id)
}

// GetAllStopSchedules retrieves all stop schedules given an id
func (scheduleStopService *ScheduleStopService) GetAllStopSchedules(context context.Context, scheduleID int64) ([]*models.ScheduleStop, error) {
	if scheduleStopService.repository == nil {
		return nil, fmt.Errorf("repositories must not be nil")
	}
	return scheduleStopService.repository.GetAll(context, scheduleID)
}

// DeleteStopSchedule deletes a stop schedule by their id
func (scheduleStopService *ScheduleStopService) DeleteStopSchedule(context context.Context, id int64) error {
	if id <= 0 {
		return nil
	}
	return scheduleStopService.repository.DeleteByID(context, id)
}

// UpdateStopSchedule updates a stop schedule by their id
func (scheduleStopService *ScheduleStopService) UpdateStopSchedule(context context.Context, id int64, updateScheduleStop *models.UpdateScheduleStop) (*models.ScheduleStop, error) {
	if id <= 0 {
		return nil, nil
	}

	if updateScheduleStop == nil {
		return nil, fmt.Errorf("update schedule stop must not be nil")
	}

	stopSchedule, err := scheduleStopService.repository.GetByID(context, id)
	if err != nil {
		return nil, fmt.Errorf("get stop schedule by id: %w", err)
	}

	if stopSchedule == nil {
		return nil, fmt.Errorf("stop schedule not found")
	}

	if updateScheduleStop.ScheduleID > 0 {
		stopSchedule.ScheduleID = updateScheduleStop.ScheduleID
	}

	if updateScheduleStop.StationID > 0 {
		stopSchedule.StationID = updateScheduleStop.StationID
	}

	if updateScheduleStop.StopOrder > 0 {
		stopSchedule.StopOrder = updateScheduleStop.StopOrder
	}

	if updateScheduleStop.ArrivalTime != "" {
		stopSchedule.ArrivalTime = updateScheduleStop.ArrivalTime
	}

	if updateScheduleStop.DepartureTime != "" {
		stopSchedule.DepartureTime = updateScheduleStop.DepartureTime
	}

	errorUpdating := scheduleStopService.repository.Update(context, stopSchedule)
	if errorUpdating != nil {
		return nil, fmt.Errorf("update stop schedule: %w", errorUpdating)
	}

	return stopSchedule, nil
}

func (scheduleStopService *ScheduleStopService) GetStopsBySchedule(context context.Context, scheduleId int64) ([]*models.ScheduleStop, error) {
	stops, err := scheduleStopService.repository.GetStopsBySchedule(context, scheduleId)
	if err != nil {
		return nil, err
	}
	return stops, nil
}
