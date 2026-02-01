package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"user-service/internal/models"
	"user-service/internal/repositories"
	"user-service/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(testing *testing.T) {
	repository, handler := setupUserTest()

	_, err := repository.Create(&models.User{
		ID:         1,
		Username:   "usernametest",
		Password:   "passwordtest",
		Email:      "emailtest",
		FiscalCode: "fiscalcodetest",
		Telephone:  "39852049548",
	})
	if err != nil {
		return
	}

	router := chi.NewRouter()
	router.Get("/users/{id}", handler.GetUser)

	request := httptest.NewRequest("GET", "/users/1", nil)
	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, request)

	assert.Equal(testing, http.StatusOK, responseRecorder.Code)
}

func TestGetUserNotFound(testing *testing.T) {
	_, handler := setupUserTest()
	router := chi.NewRouter()
	router.Get("/users/{id}", handler.GetUser)
	request := httptest.NewRequest("GET", "/users/1", nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	assert.Equal(testing, http.StatusNotFound, responseRecorder.Code)
}

func TestCreateAndGetUser(testing *testing.T) {
	_, handler := setupUserTest()

	router := chi.NewRouter()
	router.Post("/users", handler.CreateUser)
	body := `{
				"username": "usernametest", 
				"password": "passwordtest", 
				"email": "emailtest@mail.it", 
				"fiscal_code": "RSSMRA80A01H501U", 
				"telephone": "39852049548" }`
	request := httptest.NewRequest("POST", "/users", strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	assert.Equal(testing, http.StatusCreated, responseRecorder.Code)

	router.Get("/users/{id}", handler.GetUser)
	getUserRequest := httptest.NewRequest("GET", "/users/1", nil)
	newRecorder := httptest.NewRecorder()
	router.ServeHTTP(newRecorder, getUserRequest)
	assert.Equal(testing, http.StatusOK, newRecorder.Code)
}

func TestCreateUserWithInvalidBody(testing *testing.T) {
	_, handler := setupUserTest()
	router := chi.NewRouter()
	router.Post("/users", handler.CreateUser)
	body := `{"key" : "value"`
	request := httptest.NewRequest("POST", "/users", strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	assert.Equal(testing, http.StatusBadRequest, responseRecorder.Code)
}

func TestCreateUserWithBadRequest(testing *testing.T) {
	_, handler := setupUserTest()
	router := chi.NewRouter()
	router.Post("/users", handler.CreateUser)
	body := `{"key" : "value"}`
	request := httptest.NewRequest("POST", "/users", strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	assert.Equal(testing, http.StatusBadRequest, responseRecorder.Code)
}

func TestGetAllUser(testing *testing.T) {
	repository, handler := setupUserTest()

	_, err := repository.Create(&models.User{
		ID:         1,
		Username:   "usernametest",
		Password:   "passwordtest",
		Email:      "emailtest@email.it",
		FiscalCode: "VRDLGI85C60H501Z",
		Telephone:  "39852049548",
	})

	_, err = repository.Create(&models.User{
		ID:         2,
		Username:   "usernametest2",
		Password:   "passwordtest",
		Email:      "emailtest2@email.it",
		FiscalCode: "VRDLGI84C60H501Z",
		Telephone:  "39852049548",
	})

	if err != nil {
		return
	}

	router := chi.NewRouter()
	router.Get("/users", handler.GetAllUsers)
	request := httptest.NewRequest("GET", "/users", nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	assert.Equal(testing, http.StatusOK, responseRecorder.Code)
}

type FailingUserRepository struct{}

func (r *FailingUserRepository) Update(user *models.User) error {
	return errors.New("forced error")
}

func (r *FailingUserRepository) Create(user *models.User) (*models.User, error) {
	return nil, errors.New("forced error")
}

func (r *FailingUserRepository) DeleteByID(id int64) error {
	return errors.New("forced error")
}

func (r *FailingUserRepository) GetByID(id int64) (*models.User, error) {
	return nil, errors.New("forced error")
}

func (r *FailingUserRepository) GetAll() ([]*models.User, error) {
	return nil, errors.New("forced error")
}

func TestGetAllUserBadRequest(testing *testing.T) {
	failingUserRepository := &FailingUserRepository{}
	service := services.NewUserService(failingUserRepository)
	handler := NewUserHandler(service)

	router := chi.NewRouter()
	router.Get("/users", handler.GetAllUsers)
	request := httptest.NewRequest("GET", "/users", nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	assert.Equal(testing, http.StatusBadRequest, responseRecorder.Code)
}

func TestDeleteUser(testing *testing.T) {
	repository, handler := setupUserTest()
	_, _ = repository.Create(&models.User{
		Username:   "usernametest",
		Password:   "passwordtest",
		Email:      "emailtest",
		FiscalCode: "fiscalcodetest",
		Telephone:  "39852049548",
	})

	router := chi.NewRouter()
	router.Delete("/users/{id}", handler.DeleteUser)
	request := httptest.NewRequest("DELETE", "/users/1", nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	assert.Equal(testing, http.StatusNoContent, responseRecorder.Code)
}

func TestDeleteUserNotFound(testing *testing.T) {
	_, handler := setupUserTest()
	router := chi.NewRouter()
	router.Delete("/users/{id}", handler.DeleteUser)
	request := httptest.NewRequest("DELETE", "/users/1", nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	assert.Equal(testing, http.StatusNotFound, responseRecorder.Code)
}

func TestDeleteUserWithInvalidId(testing *testing.T) {
	_, handler := setupUserTest()
	router := chi.NewRouter()
	router.Delete("/users/{id}", handler.DeleteUser)
	request := httptest.NewRequest("DELETE", "/users/123abc", nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	assert.Equal(testing, http.StatusBadRequest, responseRecorder.Code)
}

func TestUpdateUser_InvalidID(t *testing.T) {
	_, handler := setupUserTest()

	router := chi.NewRouter()
	router.Patch("/users/{id}", handler.UpdateUser)

	req := httptest.NewRequest("PATCH", "/users/abc", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateUser_InvalidBody(t *testing.T) {
	_, handler := setupUserTest()

	router := chi.NewRouter()
	router.Patch("/users/{id}", handler.UpdateUser)

	body := `{"email": "missing_quote}`
	req := httptest.NewRequest("PATCH", "/users/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateUser_NotFound(t *testing.T) {
	_, handler := setupUserTest()

	router := chi.NewRouter()
	router.Patch("/users/{id}", handler.UpdateUser)

	body := `{"email": "new@mail.com"}`
	req := httptest.NewRequest("PATCH", "/users/999", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

type FailingUpdateRepository struct {
	FakeUserRepository
}

func (r *FailingUpdateRepository) Update(user *models.User) error {
	return errors.New("forced update error")
}

func TestUpdateUser_RepositoryError(t *testing.T) {
	repo := &FailingUpdateRepository{FakeUserRepository: *NewFakeUserRepository()}
	service := services.NewUserService(repo)
	handler := NewUserHandler(service)

	// Inseriamo un utente per evitare ErrUserNotFound
	repo.Create(&models.User{ID: 1, Username: "old"})

	router := chi.NewRouter()
	router.Patch("/users/{id}", handler.UpdateUser)

	body := `{"username": "newname"}`
	req := httptest.NewRequest("PATCH", "/users/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateUser_Success(t *testing.T) {
	repo, handler := setupUserTest()

	// Inseriamo un utente iniziale
	repo.Create(&models.User{
		ID:         1,
		Username:   "oldname",
		Email:      "old@mail.com",
		FiscalCode: "RSSMRA80A01H501U",
	})

	router := chi.NewRouter()
	router.Patch("/users/{id}", handler.UpdateUser)

	body := `{"username": "newname", "email": "new@mail.com"}`
	req := httptest.NewRequest("PATCH", "/users/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var updated models.User
	json.NewDecoder(rr.Body).Decode(&updated)

	assert.Equal(t, "newname", updated.Username)
	assert.Equal(t, "new@mail.com", updated.Email)
}

type FakeUserRepository struct {
	users  map[int64]*models.User
	nextID int64
}

func NewFakeUserRepository() *FakeUserRepository {
	return &FakeUserRepository{
		users:  make(map[int64]*models.User),
		nextID: 1,
	}
}

func (r *FakeUserRepository) Create(user *models.User) (*models.User, error) {
	if user.ID == 0 {
		user.ID = r.nextID
		r.nextID++
	}
	r.users[user.ID] = user
	return user, nil
}

func (r *FakeUserRepository) GetByID(id int64) (*models.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, repositories.ErrUserNotFound
	}
	return user, nil
}

func (r *FakeUserRepository) GetAll() ([]*models.User, error) {
	if len(r.users) == 0 {
		return nil, errors.New("no users found")
	}

	list := make([]*models.User, 0, len(r.users))
	for _, u := range r.users {
		list = append(list, u)
	}
	return list, nil
}

func (r *FakeUserRepository) DeleteByID(id int64) error {
	if _, ok := r.users[id]; !ok {
		return repositories.ErrUserNotFound
	}
	delete(r.users, id)
	return nil
}

func (r *FakeUserRepository) Update(user *models.User) error {
	if _, ok := r.users[user.ID]; !ok {
		return repositories.ErrUserNotFound
	}
	r.users[user.ID] = user
	return nil
}

func setupUserTest() (*FakeUserRepository, *UserHandler) {
	repository := NewFakeUserRepository()
	service := services.NewUserService(repository)
	handler := NewUserHandler(service)
	return repository, handler
}
