package models

// Schedule represents a schedule entity with its attributes.
type Schedule struct {
	ID        int     `json:"id"`
	TrainID   int     `json:"train_id:omitempty"`
	StationID int     `json:"station_id:omitempty"`
	Departure string  `json:"departure"`
	Arrival   string  `json:"arrival"`
	Status    Status  `json:"status"`
	Price     float32 `json:"price"`
}

// ScheduleStop represents a schedule stop entity with its attributes.
type ScheduleStop struct {
	ID         int    `json:"id"`
	ScheduleID int    `json:"schedule_id:omitempty"`
	StationID  int    `json:"station_id:omitempty"`
	Order      int    `json:"order"`
	Arrival    string `json:"arrival"`
	Departure  string `json:"departure"`
}

// ScheduleUpdate represents the data structure for updating a schedule.
type ScheduleUpdate struct {
	TrainID   int     `json:"train_id:omitempty"`
	StationID int     `json:"station_id:omitempty"`
	Departure string  `json:"departure"`
	Arrival   string  `json:"arrival"`
	Status    Status  `json:"status"`
	Price     float32 `json:"price"`
}

// ScheduleStopUpdate represents the data structure for updating a schedule stop.
type ScheduleStopUpdate struct {
	ScheduleID int     `json:"schedule_id:omitempty"`
	StationID  int     `json:"station_id:omitempty"`
	Order      int     `json:"order"`
	Arrival    string  `json:"arrival"`
	Departure  string  `json:"departure"`
	Price      float32 `json:"price"`
}
