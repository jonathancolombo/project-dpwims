package models

// Schedule represents a schedule entity with its attributes.
type Schedule struct {
	ID        int    `json:"id"`
	TrainID   int    `json:"train_id"`
	StationID int    `json:"station_id"`
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
	Status    string `json:"status"`
	Price     int    `json:"price"`
}

// ScheduleStop represents a schedule stop entity with its attributes.
type ScheduleStop struct {
	ID         int    `json:"id"`
	ScheduleID int    `json:"schedule_id"`
	StationID  int    `json:"station_id"`
	Order      int    `json:"order"`
	Arrival    string `json:"arrival"`
	Departure  string `json:"departure"`
}
