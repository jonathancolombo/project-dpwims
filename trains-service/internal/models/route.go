package models

// Route represents a route entity with its attributes.
type Route struct {
	ID               int64  `json:"id"`
	TrainId          string `json:"train_id"`
	DepartureStation string `json:"departure_station"`
	ArrivalStation   string `json:"arrival_station"`
	Distance         int64  `json:"distance"`
}

// UpdateRoute represents the fields that can be updated for a route entity.
type UpdateRoute struct {
	TrainId          string `json:"train_id"`
	DepartureStation string `json:"departure_station"`
	ArrivalStation   string `json:"arrival_station"`
	Distance         int64  `json:"distance"`
}
