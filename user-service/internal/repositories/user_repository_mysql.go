package repositories

import (
	"database/sql"
	"user-service/internal/models"
)

type MySQLUserRepository struct {
	database *sql.DB
}

func NewMySQLRepositoryUsers(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{database: db}
}

func (mySqlUserRepository *MySQLUserRepository) Create(user *models.User) (*models.User, error) {
	query := `INSERT INTO users (username, password, email, fiscal_code, telephone) VALUES (?, ?, ?, ?, ?)`
	result, err := mySqlUserRepository.database.Exec(query, user.Username, user.Password, user.Email, user.FiscalCode, user.Telephone)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	user.ID = id
	return user, nil
}
