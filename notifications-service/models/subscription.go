package models

// Subscription represents a user's subscription to a train's notifications.
type Subscription struct {
	ID        int64  `json:"id:omitempty"`
	UserID    string `json:"user_id:omitempty"`
	TrainUUID string `json:"train_uuid:omitempty"`
	Plan      string `json:"plan"`
}
