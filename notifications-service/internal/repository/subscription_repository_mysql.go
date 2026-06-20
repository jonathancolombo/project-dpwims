package repository

import (
	"context"
	"database/sql"
	"log"
	"notifications-service/internal/models"
)

type MySQLSubscriptionRepository struct {
	db *sql.DB
}

func NewMySQLSubscriptionRepository(db *sql.DB) *MySQLSubscriptionRepository {
	return &MySQLSubscriptionRepository{db: db}
}

func (r *MySQLSubscriptionRepository) AddSubscription(ctx context.Context, userID int64, trainUUID string, scheduleID int64) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO subscriptions (user_id, train_uuid, schedule_id) VALUES (?, ?, ?)",
		userID, trainUUID, scheduleID,
	)
	return err
}

func (r *MySQLSubscriptionRepository) GetUsersByScheduleID(ctx context.Context, scheduleID int64) ([]int64, error) {
	return r.queryUserIDs(ctx,
		"SELECT user_id FROM subscriptions WHERE schedule_id = ?",
		scheduleID,
	)
}

func (r *MySQLSubscriptionRepository) GetAllSubscriptions(ctx context.Context) ([]models.Subscription, error) {
	return r.querySubscriptions(ctx,
		"SELECT id, user_id, train_uuid, schedule_id FROM subscriptions",
	)
}

func (r *MySQLSubscriptionRepository) GetByUser(ctx context.Context, userID int64) ([]models.Subscription, error) {
	return r.querySubscriptions(ctx,
		"SELECT id, user_id, train_uuid, schedule_id FROM subscriptions WHERE user_id = ?",
		userID,
	)
}

func (r *MySQLSubscriptionRepository) GetByTrain(ctx context.Context, trainUUID string) ([]models.Subscription, error) {
	return r.querySubscriptions(ctx,
		"SELECT id, user_id, train_uuid, schedule_id FROM subscriptions WHERE train_uuid = ?",
		trainUUID,
	)
}

func (r *MySQLSubscriptionRepository) GetBySchedule(ctx context.Context, scheduleID int64) ([]models.Subscription, error) {
	return r.querySubscriptions(ctx,
		"SELECT id, user_id, train_uuid, schedule_id FROM subscriptions WHERE schedule_id = ?",
		scheduleID,
	)
}

func (r *MySQLSubscriptionRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx,
		"DELETE FROM subscriptions WHERE id = ?",
		id,
	)
	return err
}

func (r *MySQLSubscriptionRepository) querySubscriptions(ctx context.Context, query string, args ...any) ([]models.Subscription, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Println("query error:", err)
		return []models.Subscription{}, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("failed to close rows:", err)
		}
	}()

	subs := make([]models.Subscription, 0)

	for rows.Next() {
		var s models.Subscription
		if err := rows.Scan(&s.ID, &s.UserID, &s.TrainUUID, &s.ScheduleID); err != nil {
			log.Println("scan error:", err)
			return []models.Subscription{}, err
		}
		subs = append(subs, s)
	}

	return subs, nil
}

func (r *MySQLSubscriptionRepository) queryUserIDs(ctx context.Context, query string, args ...any) ([]int64, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return []int64{}, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("failed to close rows:", err)
		}
	}()

	ids := make([]int64, 0)

	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			log.Println("scan error:", err)
			return []int64{}, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}
