package repository

import (
	"context"
	"database/sql"
	"log"
	"notifications-service/internal/models"
)

// MySQLSubscriptionRepository is a concrete implementation of SubscriptionRepository that uses MySQL as the data store.
type MySQLSubscriptionRepository struct {
	database *sql.DB
}

// NewMySQLSubscriptionRepository creates a new instance of MySQLSubscriptionRepository with the given database connection.
func NewMySQLSubscriptionRepository(db *sql.DB) *MySQLSubscriptionRepository {
	return &MySQLSubscriptionRepository{database: db}
}

// AddSubscription adds a new subscription for a user to receive notifications about a specific train.
func (subscriptionRepository *MySQLSubscriptionRepository) AddSubscription(context context.Context, userID int64, trainUUID string) error {
	_, err := subscriptionRepository.database.ExecContext(context, "INSERT INTO subscriptions (user_id, train_uuid) VALUES (?, ?)", userID, trainUUID)
	return err
}

// RemoveSubscription removes an existing subscription for a user to stop receiving notifications about a specific train.
func (subscriptionRepository *MySQLSubscriptionRepository) RemoveSubscription(context context.Context, userID int64, trainUUID string) error {
	_, err := subscriptionRepository.database.ExecContext(context, "DELETE FROM subscriptions WHERE user_id = ? AND train_uuid = ?", userID, trainUUID)
	return err
}

// GetUsersByTrainUUID retrieves a list of user IDs that are subscribed to receive notifications about a specific train.
func (subscriptionRepository *MySQLSubscriptionRepository) GetUsersByTrainUUID(context context.Context, trainUUID string) ([]int64, error) {
	rows, err := subscriptionRepository.database.QueryContext(context, "SELECT user_id FROM subscriptions WHERE train_uuid = ?", trainUUID)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
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

// GetAllSubscriptions retrieves all subscriptions from the database, returning a slice of Subscription models.
func (subscriptionRepository *MySQLSubscriptionRepository) GetAllSubscriptions(context context.Context) ([]models.Subscription, error) {
	rows, err := subscriptionRepository.database.QueryContext(context, "SELECT id, user_id, train_uuid FROM subscriptions")
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("failed to close rows: ", err)
		}
	}(rows)

	var subscriptions []models.Subscription
	for rows.Next() {
		var subscription models.Subscription
		if err := rows.Scan(&subscription.ID, &subscription.UserID, &subscription.TrainUUID); err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions, nil
}

// GetByUser retrieves all subscriptions for a specific user, returning a slice of Subscription models associated with the given user ID.
func (subscriptionRepository *MySQLSubscriptionRepository) GetByUser(context context.Context, userID int64) ([]models.Subscription, error) {
	rows, err := subscriptionRepository.database.QueryContext(context, "SELECT id, user_id, train_uuid FROM subscriptions WHERE user_id = ?", userID)
	if err != nil {
		log.Println("failed to execute query: ", err)
		return []models.Subscription{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("failed to close rows: ", err)
		}
	}(rows)

	var subscriptions []models.Subscription
	for rows.Next() {
		var subscription models.Subscription
		if err := rows.Scan(&subscription.ID, &subscription.UserID, &subscription.TrainUUID); err != nil {
			log.Println("failed to scan row: ", err)
			return []models.Subscription{}, err
		}
		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions, nil
}

// GetByTrain retrieves all subscriptions for a specific train, returning a slice of Subscription models associated with the given train UUID.
func (subscriptionRepository *MySQLSubscriptionRepository) GetByTrain(context context.Context, trainUUID string) ([]models.Subscription, error) {
	rows, err := subscriptionRepository.database.QueryContext(context, "SELECT id, user_id, train_uuid FROM subscriptions WHERE train_uuid = ?", trainUUID)
	if err != nil {
		log.Println("failed to execute query: ", err)
		return []models.Subscription{}, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("failed to close rows: ", err)
		}
	}(rows)

	var subscriptions []models.Subscription
	for rows.Next() {
		var subscription models.Subscription
		if err := rows.Scan(&subscription.ID, &subscription.UserID, &subscription.TrainUUID); err != nil {
			log.Println("failed to scan row: ", err)
			return []models.Subscription{}, err
		}
		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions, nil
}

// Delete removes a subscription from the database based on its unique ID, effectively deleting the subscription record.
func (subscriptionRepository *MySQLSubscriptionRepository) Delete(context context.Context, id int64) error {
	_, err := subscriptionRepository.database.ExecContext(context, "DELETE FROM subscriptions WHERE id = ?", id)
	return err
}
