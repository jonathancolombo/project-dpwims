package repository

import (
	"context"
	"database/sql"
	"log"
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
