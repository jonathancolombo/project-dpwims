package repository

import (
	"context"
	"notifications-service/internal/models"
)

// SubscriptionRepository defines the interface for accessing subscription data related to train subscriptions.
type SubscriptionRepository interface {
	AddSubscription(context context.Context, userID int64, trainUUID string, scheduleID int64) error
	GetUsersByScheduleID(context context.Context, scheduleID int64) ([]int64, error)
	GetAllSubscriptions(context context.Context) ([]models.Subscription, error)
	GetByUser(context context.Context, userID int64) ([]models.Subscription, error)
	GetByTrain(context context.Context, trainUUID string) ([]models.Subscription, error)
	GetBySchedule(context context.Context, scheduleID int64) ([]models.Subscription, error)
	Delete(context context.Context, id int64) error
}
