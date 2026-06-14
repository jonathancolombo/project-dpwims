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
	TrainTypeRegional  = TrainType("regional")
	TrainTypeIntercity = TrainType("intercity")
	TrainTypeHighSpeed = TrainType("highspeed")
)

// Status represents the status of an entity.
type Status string

const (
	StatusActive   = Status("active")
	StatusInactive = Status("inactive")
)

// UpdateTrain represents the fields that can be updated for a train entity.
type UpdateTrain struct {
	Number   string    `json:"train_number"`
	Type     TrainType `json:"type"`
	Capacity int64     `json:"capacity"`
	Status   Status    `json:"status"`
}
