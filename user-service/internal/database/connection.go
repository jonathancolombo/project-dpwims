package database

import (
	"database/sql"
	"time"
)

const maxNumberOpenConnections = 10
const maxIdleConnections = 5

// NewMySQLConnection initializes and returns a new MySQL database connection using the provided DSN.
func NewMySQLConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxNumberOpenConnections)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxLifetime(time.Hour)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
