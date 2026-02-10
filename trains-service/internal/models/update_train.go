package models

// UpdateTrain represents the fields that can be updated for a train entity.
type UpdateTrain struct {
	Number   string `json:"train_number,omitempty"`
	Type     string `json:"type,omitempty"`
	Capacity int64  `json:"capacity,omitempty"`
	Status   string `json:"status,omitempty"`
}
