package repositories

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jonathancolombo/railway-ticketing-system/internal/models"
)

// RepositoryUsers defines the interface for a user repository
type RepositoryUsers interface {
	Create(user *models.User) (*models.User, error)
	DeleteById(user *models.User, id int64) error
	FindById(id int64) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
}

// InMemoryRepositoryUsers is an in-memory implementation of the RepositoryUsers interface
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
		Id:         newID,
		Username:   username,
		Password:   user.Password,
		FiscalCode: user.FiscalCode,
		Email:      user.Email,
		Telephone:  user.Telephone,
	}
	repository.users = append(repository.users, created)
	return created, nil
}

// DeleteById removes a user from the repository
func (repository *InMemoryRepositoryUsers) DeleteById(id int64) error {
	if id <= 0 {
		return errors.New("id must be greater than 0")
	}

	for index, user := range repository.users {
		if user.Id == id {
			repository.users = append(repository.users[:index], repository.users[index+1:]...)
			return nil
		}
	}

	return nil
}

// FindById returns a user from the repository by its ID
func (repository *InMemoryRepositoryUsers) FindById(id int64) (*models.User, error) {
	//TODO implement me
	if id <= 0 {
		return nil, errors.New("id must be greater than 0")
	}

	for _, user := range repository.users {
		if user.Id == id {
			fmt.Println("user found:")
			return user, nil
		}
	}

	fmt.Println("user not found")
	return nil, nil
}

// FindByUsername returns a user from the repository by its username
func (repository *InMemoryRepositoryUsers) FindByUsername(username string) (*models.User, error) {
	//TODO implement me
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	for _, user := range repository.users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, nil
}

// FindByEmail returns a user from the repository by its email
func (repository *InMemoryRepositoryUsers) FindByEmail(email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

// GetAll returns all users from the repository
func (repository *InMemoryRepositoryUsers) GetAll() ([]*models.User, error) {
	if len(repository.users) < 0 {
		return nil, errors.New("no users found")
	}

	return repository.users, nil
}

// NewInMemoryRepositoryUsers creates a new instance of InMemoryRepositoryUsers
func NewInMemoryRepositoryUsers() *InMemoryRepositoryUsers {
	return &InMemoryRepositoryUsers{
		users: []*models.User{
			{
				Id:         1,
				Username:   "username",
				Password:   "",
				FiscalCode: "",
				Email:      "",
				Telephone:  "",
			},
		},
	}
}
