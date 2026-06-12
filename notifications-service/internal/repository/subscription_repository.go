package repository

import (
	"context"
	"notifications-service/internal/models"
)

// SubscriptionRepository defines the interface for accessing subscription data related to train subscriptions.
type SubscriptionRepository interface {
	AddSubscription(context context.Context, userID int64, trainUUID string, plan models.Plan) error
	RemoveSubscription(context context.Context, userID int64, trainUUID string) error
	GetUsersByTrainUUID(context context.Context, trainUUID string) ([]int64, error)
	GetAllSubscriptions(context context.Context) ([]models.Subscription, error)
	GetByUser(context context.Context, userID int64) ([]models.Subscription, error)
	GetByTrain(context context.Context, trainUUID string) ([]models.Subscription, error)
	Delete(context context.Context, id int64) error
}
