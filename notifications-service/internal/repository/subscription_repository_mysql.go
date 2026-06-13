package repository

import (
	"context"
	"database/sql"
	"log"
	"notifications-service/internal/models"
)

type MySQLSubscriptionRepository struct {
	database *sql.DB
}

func NewMySQLSubscriptionRepository(db *sql.DB) *MySQLSubscriptionRepository {
	return &MySQLSubscriptionRepository{database: db}
}

func (r *MySQLSubscriptionRepository) AddSubscription(context context.Context, userID int64, trainUUID string, scheduleID int64) error {
	_, err := r.database.ExecContext(context, "INSERT INTO subscriptions (user_id, train_uuid, schedule_id) VALUES (?, ?, ?)", userID, trainUUID, scheduleID)
	return err
}

func (r *MySQLSubscriptionRepository) GetUsersByScheduleID(context context.Context, scheduleID int64) ([]int64, error) {
	rows, err := r.database.QueryContext(context, "SELECT user_id FROM subscriptions WHERE schedule_id = ?", scheduleID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Fatalf("failed to close rows: %v", err)
		}
	}(rows)

	var users []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		users = append(users, id)
	}
	return users, nil
}

func (r *MySQLSubscriptionRepository) GetAllSubscriptions(context context.Context) ([]models.Subscription, error) {
	rows, err := r.database.QueryContext(context, "SELECT id, user_id, train_uuid, schedule_id FROM subscriptions")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Println("failed to close rows: ", err)
		}
	}(rows)

	subscriptions := []models.Subscription{}
	for rows.Next() {
		var s models.Subscription
		if err := rows.Scan(&s.ID, &s.UserID, &s.TrainUUID, &s.ScheduleID); err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, s)
	}
	return subscriptions, nil
}

func (r *MySQLSubscriptionRepository) GetByUser(context context.Context, userID int64) ([]models.Subscription, error) {
	rows, err := r.database.QueryContext(context, "SELECT id, user_id, train_uuid, schedule_id FROM subscriptions WHERE user_id = ?", userID)
	if err != nil {
		log.Println("failed to execute query: ", err)
		return []models.Subscription{}, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Println("failed to close rows: ", err)
		}
	}(rows)

	subscriptions := []models.Subscription{}
	for rows.Next() {
		var s models.Subscription
		if err := rows.Scan(&s.ID, &s.UserID, &s.TrainUUID, &s.ScheduleID); err != nil {
			log.Println("failed to scan row: ", err)
			return []models.Subscription{}, err
		}
		subscriptions = append(subscriptions, s)
	}
	return subscriptions, nil
}

func (r *MySQLSubscriptionRepository) GetByTrain(context context.Context, trainUUID string) ([]models.Subscription, error) {
	rows, err := r.database.QueryContext(context, "SELECT id, user_id, train_uuid, schedule_id FROM subscriptions WHERE train_uuid = ?", trainUUID)
	if err != nil {
		log.Println("failed to execute query: ", err)
		return []models.Subscription{}, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Println("failed to close rows: ", err)
		}
	}(rows)

	subscriptions := []models.Subscription{}
	for rows.Next() {
		var s models.Subscription
		if err := rows.Scan(&s.ID, &s.UserID, &s.TrainUUID, &s.ScheduleID); err != nil {
			log.Println("failed to scan row: ", err)
			return []models.Subscription{}, err
		}
		subscriptions = append(subscriptions, s)
	}
	return subscriptions, nil
}

func (r *MySQLSubscriptionRepository) GetBySchedule(context context.Context, scheduleID int64) ([]models.Subscription, error) {
	rows, err := r.database.QueryContext(context, "SELECT id, user_id, train_uuid, schedule_id FROM subscriptions WHERE schedule_id = ?", scheduleID)
	if err != nil {
		log.Println("failed to execute query: ", err)
		return []models.Subscription{}, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Println("failed to close rows: ", err)
		}
	}(rows)

	subscriptions := []models.Subscription{}
	for rows.Next() {
		var s models.Subscription
		if err := rows.Scan(&s.ID, &s.UserID, &s.TrainUUID, &s.ScheduleID); err != nil {
			log.Println("failed to scan row: ", err)
			return []models.Subscription{}, err
		}
		subscriptions = append(subscriptions, s)
	}
	return subscriptions, nil
}

func (r *MySQLSubscriptionRepository) Delete(context context.Context, id int64) error {
	_, err := r.database.ExecContext(context, "DELETE FROM subscriptions WHERE id = ?", id)
	return err
}
