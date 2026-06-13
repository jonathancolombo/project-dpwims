package models

// Train represents a train entity with its attributes.
type Train struct {
	UUID     string    `json:"uuid"`
	Number   string    `json:"train_number"`
	Type     TrainType `json:"type"`
	Capacity int64     `json:"capacity"`
	Status   Status    `json:"status"`
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

// Status represents the status of an entity.
type Status string

const (
	// StatusActive represents an active status.
	StatusActive = Status("active")
	// StatusInactive represents an inactive status.
	StatusInactive = Status("inactive")
)

// UpdateTrain represents the fields that can be updated for a train entity.
type UpdateTrain struct {
	Number   string    `json:"train_number"`
	Type     TrainType `json:"type"`
	Capacity int64     `json:"capacity"`
	Status   Status    `json:"status"`
}
