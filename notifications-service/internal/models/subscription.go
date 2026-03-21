package models

// Subscription represents a user's subscription to a train's subscriptions.
type Subscription struct {
	ID        int64  `json:"id,omitempty"`
	UserID    int64  `json:"user_id,omitempty"`
	TrainUUID string `json:"train_uuid,omitempty"`
	Plan      Plan   `json:"plan"`
}

// Plan defines the type for different subscription plans.
type Plan string

const (
	PlanBasic   Plan = "basic"
	PlanPremium Plan = "premium"
	PlanFull    Plan = "full"
)
