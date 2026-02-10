package models

// Train represents a train entity with its attributes.
type Train struct {
	UUID     string      `json:"uuid"`
	Number   string      `json:"train_number"`
	Type     TrainType   `json:"type"`
	Capacity int64       `json:"capacity"`
	Status   TrainStatus `json:"status"`
}

// TrainType represents the type of train.
type TrainType string

const (
	// TrainTypeRegional represents a regional train type.
	TrainTypeRegional = TrainType("regional")
	// TrainTypeIntercity represents an intercity train type.
	TrainTypeIntercity = TrainType("intercity")
	// TrainTypeHighSpeed represents a high speed train type.
	TrainTypeHighSpeed = TrainType("highspeed")
)

// TrainStatus represents the status of a train.
type TrainStatus string

const (
	// StatusActive represents an active status.
	StatusActive = TrainStatus("active")
	// StatusInactive represents an inactive status.
	StatusInactive = TrainStatus("inactive")
)
