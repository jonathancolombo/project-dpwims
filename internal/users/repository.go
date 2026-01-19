package users

import (
	"errors"
	"strings"
)

type RepositoryUsers interface {
	Create(user *User) (*User, error)
	Delete(user *User) error
	FindById(id int64) (*User, error)
	FindByUsername(username string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsernameAndPassword(username, password string) (*User, error)
	GetAll() ([]*User, error)
}

type InMemoryRepositoryUsers struct {
	users []*User
}

func (repository *InMemoryRepositoryUsers) Create(user *User) (*User, error) {

	if user == nil {
		return nil, errors.New("user is nil")
	}

	newID := int64(len(repository.users) + 1)
	username := strings.ToLower(user.Username)
	created := &User{
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

func (repository *InMemoryRepositoryUsers) Delete(user *User) error {
	//TODO implement me
	panic("implement me")
}

func (repository *InMemoryRepositoryUsers) FindById(id int64) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *InMemoryRepositoryUsers) FindByUsername(username string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *InMemoryRepositoryUsers) FindByEmail(email string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *InMemoryRepositoryUsers) FindByUsernameAndPassword(username, password string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *InMemoryRepositoryUsers) GetAll() ([]*User, error) {
	return repository.users, nil
}

func NewInMemoryRepositoryUsers() RepositoryUsers {
	return &InMemoryRepositoryUsers{
		users: []*User{
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
