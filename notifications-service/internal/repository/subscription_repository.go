package repository

import "context"

// SubscriptionRepository defines the interface for accessing subscription data related to train notifications.
type SubscriptionRepository interface {
	AddSubscription(context context.Context, userID int64, trainUUID string) error
	RemoveSubscription(context context.Context, userID int64, trainUUID string) error
	GetUsersByTrainUUID(context context.Context, trainUUID string) ([]int64, error)
}
