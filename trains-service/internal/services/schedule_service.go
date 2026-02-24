package services

import (
	"context"
	"fmt"
	"trains-service/internal/models"
	"trains-service/internal/repositories"
)

// ScheduleService defines the interface for managing Schedule entities.
type ScheduleService struct {
	repository repositories.IScheduleRepository
}

// NewScheduleService creates a new ScheduleService instance
func NewScheduleService(repository repositories.IScheduleRepository) *ScheduleService {
	return &ScheduleService{
		repository: repository,
	}
}

// CreateSchedule creates a new schedule
func (scheduleService *ScheduleService) CreateSchedule(context context.Context, schedule *models.Schedule) (*models.Schedule, error) {
	if schedule == nil {
		return nil, fmt.Errorf("schedule must not be nil")
	}

	if schedule.ID <= 0 {
		return nil, fmt.Errorf("route id must be greater than 0")
	}
	if schedule.TrainID == "" {
		return nil, fmt.Errorf("train id must be greater than 0")
	}

	if schedule.StationID <= 0 {
		return nil, fmt.Errorf("station id must be greater than 0")
	}

	if schedule.Departure == "" {
		return nil, fmt.Errorf("departure must not be empty")
	}

	if schedule.Arrival == "" {
		return nil, fmt.Errorf("arrival must not be empty")
	}

	if schedule.Price < 0 {
		return nil, fmt.Errorf("price must be greater than or equal to zero")
	}

	if schedule.Status == "" {
		return nil, fmt.Errorf("status must not be empty")
	}

	return scheduleService.repository.Create(context, schedule)

}

// GetSchedule retrieves a schedule by their id
func (scheduleService *ScheduleService) GetSchedule(context context.Context, id int64) (*models.Schedule, error) {
	if id <= 0 {
		return nil, fmt.Errorf("id must be greater than 0")
	}
	return scheduleService.repository.GetByID(context, id)
}

// GetAllSchedules retrieves all schedules
func (scheduleService *ScheduleService) GetAllSchedules(context context.Context) ([]*models.Schedule, error) {
	if scheduleService.repository == nil {
		return nil, fmt.Errorf("repositories must not be nil")
	}
	return scheduleService.repository.GetAll(context)
}

// DeleteSchedule deletes a schedule by their id
func (scheduleService *ScheduleService) DeleteSchedule(context context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("id must be greater than 0")
	}
	return scheduleService.repository.DeleteByID(context, id)
}

// UpdateSchedule updates a schedule by their id
func (scheduleService *ScheduleService) UpdateSchedule(context context.Context, id int64, updateSchedule *models.UpdateSchedule) (*models.Schedule, error) {
	if id <= 0 {
		return nil, fmt.Errorf("id must be greater than 0")
	}

	if updateSchedule == nil {
		return nil, fmt.Errorf("update schedule must not be nil")
	}

	schedule, err := scheduleService.repository.GetByID(context, id)
	if err != nil {
		return nil, fmt.Errorf("get schedule by id: %w", err)
	}

	if schedule == nil {
		return nil, fmt.Errorf("schedule must not be nil")
	}

	if updateSchedule.TrainID != "" {
		schedule.TrainID = updateSchedule.TrainID
	}

	if updateSchedule.StationID > 0 {
		schedule.StationID = updateSchedule.StationID
	}

	if updateSchedule.Departure != "" {
		schedule.Departure = updateSchedule.Departure
	}

	if updateSchedule.Arrival != "" {
		schedule.Arrival = updateSchedule.Arrival
	}

	if updateSchedule.Price >= 0 {
		schedule.Price = updateSchedule.Price
	}

	if updateSchedule.Status != "" && (schedule.Status == models.StatusActive || schedule.Status == models.StatusInactive) {
		schedule.Status = updateSchedule.Status
	}

	errorUpdating := scheduleService.repository.Update(context, schedule)
	if errorUpdating != nil {
		return nil, fmt.Errorf("update schedule: %w", errorUpdating)
	}

	return schedule, nil
}
