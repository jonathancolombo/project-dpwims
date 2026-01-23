package repositories

import (
	"errors"
	"sync"
	"user-service/internal/models"
)

type InMemoryUserRepository struct {
	rwMutex sync.RWMutex
	data    map[int64]*models.User
	nextID  int64
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		data:   make(map[int64]*models.User),
		nextID: 1,
	}
}

// Create a method to create users into memory
func (inMemoryUserRepository *InMemoryUserRepository) Create(user *models.User) (*models.User, error) {
	inMemoryUserRepository.rwMutex.Lock()
	defer inMemoryUserRepository.rwMutex.Unlock()

	for _, currentUser := range inMemoryUserRepository.data {
		if currentUser.Username == user.Username {
			return nil, errors.New("user already exists")
		}
	}

	user.ID = inMemoryUserRepository.nextID
	inMemoryUserRepository.nextID++
	inMemoryUserRepository.data[user.ID] = user
	return user, nil
}

// FindByID is a method that search a user by his id and returns it
func (inMemoryUserRepository *InMemoryUserRepository) FindByID(id int64) (*models.User, error) {
	inMemoryUserRepository.rwMutex.RLock()
	defer inMemoryUserRepository.rwMutex.RUnlock()

	user, exists := inMemoryUserRepository.data[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}

func (inMemoryUserRepository *InMemoryUserRepository) FindByUsername(username string) (*models.User, error) {
	inMemoryUserRepository.rwMutex.RLock()
	defer inMemoryUserRepository.rwMutex.RUnlock()

	for _, user := range inMemoryUserRepository.data {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, nil
}

// DeleteByID is a method that's delete a user by id into memory
func (inMemoryUserRepository *InMemoryUserRepository) DeleteByID(id int64) error {
	inMemoryUserRepository.rwMutex.Lock()
	defer inMemoryUserRepository.rwMutex.Unlock()

	if _, exists := inMemoryUserRepository.data[id]; !exists {
		return errors.New("user not found")
	}

	delete(inMemoryUserRepository.data, id)
	return nil
}

func (inMemoryUserRepository *InMemoryUserRepository) GetAll() ([]*models.User, error) {
	inMemoryUserRepository.rwMutex.RLock()
	defer inMemoryUserRepository.rwMutex.RUnlock()

	users := make([]*models.User, 0, len(inMemoryUserRepository.data))
	for _, user := range inMemoryUserRepository.data {
		users = append(users, user)
	}
	return users, nil
}
