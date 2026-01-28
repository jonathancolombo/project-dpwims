package repositories

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"user-service/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMySQLCreateUserSuccessfully(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	if err != nil {
		testing.Fatalf("failed to create sqlMock: %v", err)
	}

	defer func(databaseMock *sql.DB) {
		err := databaseMock.Close()
		if err != nil {

		}
	}(databaseMock)

	repo := NewMySQLRepositoryUsers(databaseMock)
	user := &models.User{
		Username:     "testuser",
		Password:     "password123",
		PasswordSalt: "16",
		Email:        "test@example.com",
		FiscalCode:   "ABC123",
		Telephone:    "1234567890",
	}

	query := "INSERT INTO users (username, password, password_salt, email, fiscal_code, telephone) VALUES (?, ?, ?, ?, ?, ?)"

	sqlMock.ExpectPrepare(regexp.QuoteMeta(query)).
		ExpectExec().
		WithArgs("testuser", "password123", "16", "test@example.com", "ABC123", "1234567890").
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	result, err := repo.Create(ctx, user)

	assert.NoError(testing, err)
	assert.Equal(testing, int64(1), result.ID)
	assert.Equal(testing, "testuser", result.Username)

	assert.NoError(testing, sqlMock.ExpectationsWereMet())
}

func TestMySQLCreateUserNilUser(testing *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		testing.Fatalf("failed to create mock: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	repository := NewMySQLRepositoryUsers(db)
	contextBG := context.Background()

	result, err := repository.Create(contextBG, nil)

	assert.Nil(testing, result)
	assert.Error(testing, err)
}
