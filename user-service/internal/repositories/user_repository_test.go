package repositories

import (
	"errors"
	"testing"
	"user-service/internal/models"

	"github.com/stretchr/testify/assert"
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
