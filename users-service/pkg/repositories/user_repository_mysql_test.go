package repositories

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"users-service/pkg/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMySQLCreateUserSuccessfully(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(testing, err)
	if err != nil {
		testing.Fatalf("failed to create sqlMock: %v", err)
	}

	defer func(databaseMock *sql.DB) {
		err := databaseMock.Close()
		if err != nil {

		}
	}(databaseMock)

	repository := NewMySQLRepositoryUsers(databaseMock)
	user := &models.User{
		Username:     "testuser",
		Password:     "password123",
		PasswordSalt: "16",
		Email:        "test@example.com",
		FiscalCode:   "ABC123",
		Telephone:    "1234567890",
		Role:         models.RoleAdmin,
	}

	query := "INSERT INTO users (username, password, password_salt, email, fiscal_code, telephone, role) VALUES (?, ?, ?, ?, ?, ?, ?)"

	sqlMock.ExpectPrepare(regexp.QuoteMeta(query)).
		ExpectExec().
		WithArgs("testuser", "password123", "16", "test@example.com", "ABC123", "1234567890", models.RoleAdmin).
		WillReturnResult(sqlmock.NewResult(1, 1))

	emptyContext := context.Background()
	result, err := repository.Create(emptyContext, user)

	assert.NoError(testing, err)
	assert.Equal(testing, int64(1), result.ID)
	assert.Equal(testing, "testuser", result.Username)
	assert.NoError(testing, sqlMock.ExpectationsWereMet())
}

func TestMySQLCreateUserNilUser(testing *testing.T) {
	database, _, err := sqlmock.New()
	require.NoError(testing, err)
	if err != nil {
		testing.Fatalf("failed to create mock: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(database)

	repository := NewMySQLRepositoryUsers(database)
	emptyContext := context.Background()

	result, err := repository.Create(emptyContext, nil)

	assert.Nil(testing, result)
	assert.Error(testing, err)
}

func TestFindByIDSuccessfully(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(testing, err)
	if err != nil {
		testing.Fatalf("failed to create sqlMock: %v", err)
	}

	repository := NewMySQLRepositoryUsers(databaseMock)
	expectedID := int64(1)
	expectedUser := models.User{
		ID:         expectedID,
		Username:   "username",
		Password:   "passowordhashed",
		Email:      "username@email.com",
		FiscalCode: "ABCDEF12G34H567I",
		Telephone:  "1234567890",
		Role:       models.RoleCustomer,
	}

	rows := sqlMock.NewRows([]string{"id", "username", "password", "email", "fiscal_code", "telephone", "role"}).AddRow(expectedUser.ID, expectedUser.Username, expectedUser.Password, expectedUser.Email, expectedUser.FiscalCode, expectedUser.Telephone, expectedUser.Role)
	sqlMock.ExpectQuery("SELECT id, username, password, email, fiscal_code, telephone, role FROM users WHERE id = ?").WithArgs(expectedID).WillReturnRows(rows)
	user, err := repository.GetByID(context.Background(), expectedID)
	require.NoError(testing, err)
	require.NotNil(testing, user)
	assert.Equal(testing, expectedUser, *user)
}

func TestFindByID_NotFound(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(testing, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(databaseMock)

	repository := &MySQLUserRepository{database: databaseMock}

	sqlMock.ExpectQuery("SELECT id, username, password, email, fiscal_code, telephone, role FROM users WHERE id = ?").
		WithArgs(int64(99)).
		WillReturnError(sql.ErrNoRows)

	user, err := repository.GetByID(context.Background(), 99)

	require.Error(testing, err)
	assert.Nil(testing, user)
	assert.ErrorIs(testing, err, ErrUserNotFound)
}

func TestFindByID_ScanError(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(testing, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(databaseMock)

	repository := &MySQLUserRepository{database: databaseMock}

	sqlMock.ExpectQuery("SELECT id, username, password, email, fiscal_code, telephone, role FROM users WHERE id = ?").
		WithArgs(int64(1)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id", "username", "password", "email", "fiscal_code", "telephone", "role",
			}).AddRow(
				"invalid-int",
				"john",
				"pass",
				"john@mail.com",
				"ABCDEF12G34H567I",
				"1234567890",
				models.RoleAdmin,
			),
		)

	user, err := repository.GetByID(context.Background(), 1)

	require.Error(testing, err)
	assert.Nil(testing, user)
	assert.Contains(testing, err.Error(), "scan user")
}

func TestGetAll_Success(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(testing, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(databaseMock)

	repo := &MySQLUserRepository{database: databaseMock}

	rows := sqlmock.NewRows([]string{
		"id", "username", "password", "email", "fiscal_code", "telephone", "role",
	}).
		AddRow(1, "john", "hashed1", "john@mail.com", "ABCDEF12G34H567I", "1234567890", models.RoleAdmin).
		AddRow(2, "mary", "hashed2", "mary@mail.com", "LMNOPQ12R34S567T", "0987654321", models.RoleCustomer)

	sqlMock.ExpectQuery("SELECT id, username, password, email, fiscal_code, telephone, role FROM users").
		WillReturnRows(rows)

	users, err := repo.GetAll(context.Background())

	require.NoError(testing, err)
	require.Len(testing, users, 2)

	assert.Equal(testing, int64(1), users[0].ID)
	assert.Equal(testing, "john", users[0].Username)

	assert.Equal(testing, int64(2), users[1].ID)
	assert.Equal(testing, "mary", users[1].Username)
}

func TestGetAll_Empty(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(testing, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(databaseMock)

	repository := &MySQLUserRepository{database: databaseMock}

	rows := sqlmock.NewRows([]string{
		"id", "username", "password", "email", "fiscal_code", "telephone",
	})

	sqlMock.ExpectQuery("SELECT id, username, password, email, fiscal_code, telephone, role FROM users").
		WillReturnRows(rows)

	users, err := repository.GetAll(context.Background())

	require.NoError(testing, err)
	require.Empty(testing, users)
}

func TestGetAll_QueryError(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(testing, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(databaseMock)

	repository := &MySQLUserRepository{database: databaseMock}

	sqlMock.ExpectQuery("SELECT id, username, password, email, fiscal_code, telephone, role FROM users").
		WillReturnError(errors.New("databaseMock failure"))

	users, err := repository.GetAll(context.Background())

	require.Error(testing, err)
	assert.Nil(testing, users)
	assert.Contains(testing, err.Error(), "query all users")
}

func TestGetAll_ScanError(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(testing, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(databaseMock)

	repository := &MySQLUserRepository{database: databaseMock}

	rows := sqlmock.NewRows([]string{
		"id", "username", "password", "email", "fiscal_code", "telephone", "role",
	}).AddRow(
		"invalid-int",
		"john",
		"hashed",
		"john@mail.com",
		"ABCDEF12G34H567I",
		"1234567890",
		models.RoleAdmin,
	)

	sqlMock.ExpectQuery("SELECT id, username, password, email, fiscal_code, telephone, role FROM users").
		WillReturnRows(rows)

	users, err := repository.GetAll(context.Background())

	require.Error(testing, err)
	assert.Nil(testing, users)
	assert.Contains(testing, err.Error(), "scan user")
}

func TestDeleteByID_Success(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(testing, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(databaseMock)

	repository := &MySQLUserRepository{database: databaseMock}

	sqlMock.ExpectExec("DELETE FROM users WHERE id = ?").
		WithArgs(int64(1)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repository.DeleteByID(context.Background(), 1)

	require.NoError(testing, err)
}

func TestDeleteByID_NotFound(t *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(databaseMock)

	repository := &MySQLUserRepository{database: databaseMock}

	sqlMock.ExpectExec("DELETE FROM users WHERE id = ?").
		WithArgs(int64(99)).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = repository.DeleteByID(context.Background(), 99)

	require.Error(t, err)
	assert.ErrorIs(t, err, ErrUserNotFound)
}

func TestDeleteByID_QueryError(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(testing, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(databaseMock)

	repository := &MySQLUserRepository{database: databaseMock}

	sqlMock.ExpectExec("DELETE FROM users WHERE id = ?").
		WithArgs(int64(1)).
		WillReturnError(errors.New("databaseMock failure"))

	err = repository.DeleteByID(context.Background(), 1)

	require.Error(testing, err)
	assert.Contains(testing, err.Error(), "delete user")
}

func TestDeleteByID_RowsAffectedError(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(testing, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(databaseMock)

	repository := &MySQLUserRepository{database: databaseMock}

	sqlMock.ExpectExec("DELETE FROM users WHERE id = ?").
		WithArgs(int64(1)).
		WillReturnResult(sqlmock.NewErrorResult(errors.New("rows error")))

	err = repository.DeleteByID(context.Background(), 1)

	require.Error(testing, err)
	assert.Contains(testing, err.Error(), "rows affected")
}

func TestUpdateByID_Success(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(testing, err)
	defer func(databaseMock *sql.DB) {
		err := databaseMock.Close()
		if err != nil {

		}
	}(databaseMock)

	repository := &MySQLUserRepository{database: databaseMock}
	user := &models.User{
		ID:           int64(1),
		Username:     "john",
		Password:     "abcefghijklmn",
		Email:        "john@mail.it",
		FiscalCode:   "MCJDBNFVKCVSDÈG",
		Telephone:    "25645456",
		Role:         models.RoleAdmin,
		PasswordSalt: "16",
	}

	sqlMock.ExpectExec(`UPDATE users`).
		WithArgs(
			user.Username,
			user.Email,
			user.Telephone,
			user.FiscalCode,
			user.Role,
			user.Password,
			user.PasswordSalt,
			user.ID,
		).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repository.Update(context.Background(), user)
	require.NoError(testing, err)
	require.NoError(testing, sqlMock.ExpectationsWereMet())
}

func TestUpdateByID_NotFound(testing *testing.T) {
	databaseMock, sqlMock, err := sqlmock.New()
	require.NoError(testing, err)
	defer func(databaseMock *sql.DB) {
		err := databaseMock.Close()
		if err != nil {

		}
	}(databaseMock)

	repository := &MySQLUserRepository{database: databaseMock}
	user := &models.User{
		ID:           int64(1),
		Username:     "john",
		Password:     "abcefghijklmn",
		Email:        "john@mail.it",
		FiscalCode:   "MCJDBNFVKCVSDÈG",
		Telephone:    "25645456",
		Role:         models.RoleCustomer,
		PasswordSalt: "16",
	}

	sqlMock.ExpectExec(`UPDATE users`).
		WithArgs(
			user.Username,
			user.Email,
			user.Telephone,
			user.FiscalCode,
			user.Role,
			user.Password,
			user.PasswordSalt,
			user.ID,
		).
		WillReturnError(errors.New("db error"))

	err = repository.Update(context.Background(), user)
	require.Error(testing, err)
	assert.Contains(testing, err.Error(), "db error")
	require.NoError(testing, sqlMock.ExpectationsWereMet())
}
