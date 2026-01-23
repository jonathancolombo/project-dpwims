package repositories

import (
	"errors"
	"sync"
	"user-service/internal/models"
)

type InMemoryUserRepository struct {
	mu     sync.RWMutex
	data   map[int64]*models.User
	nextID int64
}

func (r *InMemoryUserRepository) FindById(id int64) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		data:   make(map[int64]*models.User),
		nextID: 1,
	}
}

func (r *InMemoryUserRepository) Create(user *models.User) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.nextID++

	r.data[user.ID] = user
	return user, nil
}

func (r *InMemoryUserRepository) FindByID(id int64) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.data[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}

func (r *InMemoryUserRepository) FindByUsername(username string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.data {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, nil
}

func (r *InMemoryUserRepository) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.data, id)
	return nil
}

func (r *InMemoryUserRepository) GetAll() ([]*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*models.User, 0, len(r.data))
	for _, user := range r.data {
		users = append(users, user)
	}
	return users, nil
}
