package repositories

import (
	"errors"
	"testing"
	"user-service/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUserSuccessfully(testing *testing.T) {
	var repository = NewInMemoryRepositoryUsers()
	var user = &models.User{
		Username:   "UsernameTest",
		Password:   "passwordtest",
		Email:      "emailtest",
		FiscalCode: "fiscalcodetest",
		Telephone:  "39852049548",
	}

	userCreated, err := repository.Create(user)
	if err != nil {
		return
	}
	assert.Equal(testing, int64(1), userCreated.ID)
	assert.Equal(testing, "usernametest", userCreated.Username)
	assert.Equal(testing, "emailtest", userCreated.Email)
	assert.Equal(testing, "fiscalcodetest", userCreated.FiscalCode)
	assert.Equal(testing, "39852049548", userCreated.Telephone)
	assert.Equal(testing, 1, len(repository.users))
}

func TestCreateUserNil(testing *testing.T) {
	repository := NewInMemoryRepositoryUsers()
	userCreated, err := repository.Create(nil)
	assert.Nil(testing, userCreated)
	assert.Error(testing, err)
}

func TestCreateUserDeleteByIdSuccessfully(testing *testing.T) {
	var repository = NewInMemoryRepositoryUsers()

	var user = &models.User{
		Username:   "UsernameTest",
		Password:   "passwordtest",
		Email:      "emailtest",
		FiscalCode: "fiscalcodetest",
		Telephone:  "39852049548",
	}
	_, err := repository.Create(user)
	var usefulIdUser int64 = 1
	var result = repository.DeleteByID(usefulIdUser)
	assert.Nil(testing, result)
	assert.NoError(testing, err)
}

func TestCreateUserDeleteByIdWithIdMinorOrEqualZero(testing *testing.T) {
	repository := NewInMemoryRepositoryUsers()
	var resultIdZero = repository.DeleteByID(int64(0))
	var err = errors.New("id must be greater than 0")
	var resultMinorZero = repository.DeleteByID(int64(-1))
	assert.Error(testing, err, resultIdZero)
	assert.Error(testing, err, resultMinorZero)
}

func TestDeleteByIdWithIdNotFound(testing *testing.T) {
	var repository = NewInMemoryRepositoryUsers()
	var usefulIdUser int64 = 1
	var result = repository.DeleteByID(usefulIdUser)
	assert.Nil(testing, result)
}

func TestGetByID_InvalidID(t *testing.T) {
	repo := &InMemoryRepositoryUsers{
		users: []*models.User{}, // non importa, fallisce prima
	}

	// Supponiamo che checkIdIntoUsersList(0) ritorni errore
	user, err := repo.GetByID(0)

	require.Error(t, err)
	assert.Nil(t, user)
}

func TestGetByID_UserFound(t *testing.T) {
	expected := &models.User{
		ID:         1,
		Username:   "john",
		Email:      "john@mail.com",
		FiscalCode: "ABCDEF12G34H567I",
		Telephone:  "1234567890",
	}

	repo := &InMemoryRepositoryUsers{
		users: []*models.User{
			expected,
			{ID: 2, Username: "mary"},
		},
	}

	user, err := repo.GetByID(1)

	require.NoError(t, err)
	require.NotNil(t, user)
	assert.Equal(t, expected, user)
}

func TestGetByID_UserNotFound(t *testing.T) {
	repo := &InMemoryRepositoryUsers{
		users: []*models.User{
			{ID: 1, Username: "john"},
		},
	}

	user, err := repo.GetByID(99)

	require.NoError(t, err)
	assert.Nil(t, user)
}

func TestGetByID_EmptyList(t *testing.T) {
	repo := &InMemoryRepositoryUsers{
		users: []*models.User{},
	}

	user, err := repo.GetByID(1)

	require.NoError(t, err)
	assert.Nil(t, user)
}

func TestInMemoryRepositoryUsers_GetAll_Success(t *testing.T) {
	repo := &InMemoryRepositoryUsers{
		users: []*models.User{
			{ID: 1, Username: "john"},
			{ID: 2, Username: "mary"},
		},
	}

	users, err := repo.GetAll()

	require.NoError(t, err)
	require.Len(t, users, 2)
	assert.Equal(t, int64(1), users[0].ID)
	assert.Equal(t, "john", users[0].Username)
	assert.Equal(t, int64(2), users[1].ID)
}

func TestInMemoryRepositoryUsers_GetAll_NoUsers(t *testing.T) {
	repo := &InMemoryRepositoryUsers{
		users: []*models.User{}, // lista vuota
	}

	users, err := repo.GetAll()

	require.Error(t, err)
	assert.Nil(t, users)
	assert.EqualError(t, err, "no users found")
}

func TestInMemoryRepositoryUsers_GetAll_NilSlice(t *testing.T) {
	repo := &InMemoryRepositoryUsers{
		users: nil,
	}

	users, err := repo.GetAll()

	require.Error(t, err)
	assert.Nil(t, users)
	assert.EqualError(t, err, "no users found")
}
