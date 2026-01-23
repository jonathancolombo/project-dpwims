package repositories

import (
	"errors"
	"fmt"
	"strings"
	"user-service/internal/models"
)

// IUserRepository defines the interface for a user repository
type IUserRepository interface {
	Create(user *models.User) (*models.User, error)
	DeleteByID(id int64) error
	FindByID(id int64) (*models.User, error)
	GetAll() ([]*models.User, error)
}

// InMemoryRepositoryUsers is an in-memory implementation of the IUserRepository interface
type InMemoryRepositoryUsers struct {
	users []*models.User
}

// Create adds a new user to the repository
func (repository *InMemoryRepositoryUsers) Create(user *models.User) (*models.User, error) {

	if user == nil {
		return nil, errors.New("user is nil")
	}

	newID := int64(len(repository.users) + 1)
	username := strings.ToLower(user.Username)

	created := &models.User{
		ID:         newID,
		Username:   username,
		Password:   user.Password,
		FiscalCode: user.FiscalCode,
		Email:      user.Email,
		Telephone:  user.Telephone,
	}
	repository.users = append(repository.users, created)
	return created, nil
}

// DeleteByID removes a user from the repository
func (repository *InMemoryRepositoryUsers) DeleteByID(id int64) error {
	if id <= 0 {
		return errors.New("id must be greater than 0")
	}

	for index, user := range repository.users {
		if user.ID == id {
			repository.users = append(repository.users[:index], repository.users[index+1:]...)
			return nil
		}
	}

	return nil
}

// FindByID returns a user from the repository by its ID
func (repository *InMemoryRepositoryUsers) FindByID(id int64) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("id must be greater than 0")
	}

	for _, user := range repository.users {
		if user.ID == id {
			fmt.Println("user found:")
			return user, nil
		}
	}

	fmt.Println("user not found")
	return nil, nil
}

// GetAll returns all users from the repository
func (repository *InMemoryRepositoryUsers) GetAll() ([]*models.User, error) {
	if len(repository.users) <= 0 {
		return nil, errors.New("no users found")
	}

	return repository.users, nil
}

// NewInMemoryRepositoryUsers creates a new instance of InMemoryRepositoryUsers
func NewInMemoryRepositoryUsers() *InMemoryRepositoryUsers {
	return &InMemoryRepositoryUsers{
		users: []*models.User{
			{
				ID:         0,
				Username:   "",
				Password:   "",
				FiscalCode: "",
				Email:      "",
				Telephone:  "",
			},
		},
	}
}
