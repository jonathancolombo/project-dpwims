package services

import (
	"errors"
	"strings"
	"user-service/internal/models"
	"user-service/internal/repositories"
)

type UserService struct {
	repository repositories.IUserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(repository repositories.IUserRepository) *UserService {
	return &UserService{repository: repository}
}

// CreateUser creates a new user
func (service *UserService) CreateUser(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}
	if strings.TrimSpace(user.Username) == "" {
		return nil, errors.New("username is required")
	}

	user.Username = strings.ToLower(strings.TrimSpace(user.Username))
	return service.repository.Create(user)
}

// GetAllUsers retrieves all users
func (service *UserService) GetAllUsers() ([]*models.User, error) {
	if service.repository == nil {
		return nil, errors.New("repository is nil")
	}
	return service.repository.GetAll()
}

// DeleteUserByID deletes a user by their ID
func (service *UserService) DeleteUserByID(id int64) error {
	return service.repository.DeleteByID(id)
}

func (service *UserService) GetUser(id int64) (*models.User, error) {
	return service.repository.GetByID(id)
}
