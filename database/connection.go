package database

import (
	"database/sql"
	"time"
)

const maxNumberOpenConnections = 10
const maxIdleConnections = 5

// NewMySQLConnection initializes and returns a new MySQL database connection using the provided DSN.
func NewMySQLConnection(dsn string) (*sql.DB, error) {
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	database.SetMaxOpenConns(maxNumberOpenConnections)
	database.SetMaxIdleConns(maxIdleConnections)
	database.SetConnMaxLifetime(time.Hour)
	if err := database.Ping(); err != nil {
		return nil, err
	}
	return database, nil
}
