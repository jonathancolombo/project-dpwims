package services_test

import (
	"context"
	"testing"
	"users-service/internal/models"
	"users-service/internal/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) DeleteByID(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	//argUser = args.Get(0).(*models.User)
	return args.Error(0)
}

func (m *MockUserRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.User), args.Error(1)
}

func TestUserService_CreateUser(testing *testing.T) {
	repository := new(MockUserRepository)
	userService := services.NewUserService(repository)

	user := &models.User{
		ID:         0,
		Username:   "username",
		Password:   "Password123!",
		Email:      "username@mail.com",
		FiscalCode: "BNCLNZ90B15F205X",
		Telephone:  "321654864651",
		Role:       "user",
	}
	repository.On("Create", mock.Anything, mock.AnythingOfType("*models.User")).Return(user, nil)
	result, err := userService.CreateUser(context.Background(), user)
	assert.NoError(testing, err)
	assert.NotNil(testing, result)
	assert.Equal(testing, "username", result.Username)
	assert.Equal(testing, "username@mail.com", result.Email)
	assert.Equal(testing, "BNCLNZ90B15F205X", result.FiscalCode)
	assert.NotEqual(testing, "Password123!", result.Password)
	assert.NotEmpty(testing, result.PasswordSalt)
	assert.Len(testing, result.PasswordSalt, 32) // Ruolo
	assert.Equal(testing, "user", result.Role)
	repository.AssertExpectations(testing)
}

func TestUpdateUser_Success(t *testing.T) {
	repo := new(MockUserRepository)
	service := services.NewUserService(repo)

	existingUser := &models.User{
		ID:         1,
		Username:   "oldname",
		Email:      "old@example.com",
		Telephone:  "1111111111",
		FiscalCode: "OLDOLD80A01H501U",
		Role:       "user",
		Password:   "oldhash",
	}

	repo.On("GetByID", mock.Anything, int64(1)).Return(existingUser, nil)

	repo.On("Update", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

	newEmail := "new@example.com"
	newUsername := "newname"

	req := models.UpdateUserRequest{
		Email:    &newEmail,
		Username: &newUsername,
	}

	updated, err := service.UpdateUser(context.Background(), 1, req)

	assert.NoError(t, err)
	assert.Equal(t, "newname", updated.Username)
	assert.Equal(t, "new@example.com", updated.Email)

	repo.AssertExpectations(t)
}
