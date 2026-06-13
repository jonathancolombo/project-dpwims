package models

// Station represents a train station entity with its associated fields.
type Station struct {
	ID     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	City   string `json:"city,omitempty"`
	Region string `json:"region,omitempty"`
	Status Status `json:"status,omitempty"`
}

// UpdateStation represents the fields that can be updated for a station entity.
type UpdateStation struct {
	Name   string `json:"name,omitempty"`
	City   string `json:"city,omitempty"`
	Region string `json:"region,omitempty"`
	Status Status `json:"status,omitempty"`
}
