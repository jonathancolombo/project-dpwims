package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"users-service/pkg/models"
	"users-service/pkg/repositories"
	"users-service/pkg/services"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

type FakeUserRepository struct {
	users  map[int64]*models.User
	nextID int64
}

func (r *FakeUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	panic("implement me")
}

func NewFakeUserRepository() *FakeUserRepository {
	return &FakeUserRepository{
		users:  make(map[int64]*models.User),
		nextID: 1,
	}
}

func (r *FakeUserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if user.ID == 0 {
		user.ID = r.nextID
		r.nextID++
	}
	r.users[user.ID] = user
	return user, nil
}

func (r *FakeUserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, repositories.ErrUserNotFound
	}
	return user, nil
}

func (r *FakeUserRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	if len(r.users) == 0 {
		return nil, errors.New("no users found")
	}

	list := make([]*models.User, 0, len(r.users))
	for _, u := range r.users {
		list = append(list, u)
	}
	return list, nil
}

func (r *FakeUserRepository) DeleteByID(ctx context.Context, id int64) error {
	if _, ok := r.users[id]; !ok {
		return repositories.ErrUserNotFound
	}
	delete(r.users, id)
	return nil
}

func (r *FakeUserRepository) Update(ctx context.Context, user *models.User) error {
	if _, ok := r.users[user.ID]; !ok {
		return repositories.ErrUserNotFound
	}
	r.users[user.ID] = user
	return nil
}

func TestGetUser(testing *testing.T) {
	repository, handler := setupUserTest()
	emptyContext := context.Background()
	_, err := repository.Create(emptyContext, &models.User{
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
				"telephone": "39852049548" ,
				"role" : "admin"
				}`

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
	emptyContext := context.Background()

	_, err := repository.Create(emptyContext, &models.User{
		ID:         1,
		Username:   "usernametest",
		Password:   "passwordtest",
		Email:      "emailtest@email.it",
		FiscalCode: "VRDLGI85C60H501Z",
		Telephone:  "39852049548",
	})

	_, err = repository.Create(emptyContext, &models.User{
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

func TestGetAllUserBadRequest(testing *testing.T) {
	failingUserRepository := &FakeUserRepository{}
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
	emptyContext := context.Background()
	_, _ = repository.Create(emptyContext, &models.User{
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

func (r *FailingUpdateRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *FailingUpdateRepository) Update(context context.Context, user *models.User) error {
	return errors.New("forced update error")
}

func TestUpdateUser_RepositoryError(testing *testing.T) {
	repo := &FailingUpdateRepository{FakeUserRepository: *NewFakeUserRepository()}
	service := services.NewUserService(repo)
	handler := NewUserHandler(service)
	emptyContext := context.Background()
	_, err := repo.Create(emptyContext, &models.User{ID: 1, Username: "old"})
	if err != nil {
		return
	}

	router := chi.NewRouter()
	router.Patch("/users/{id}", handler.UpdateUser)

	body := `{"username": "newname"}`
	req := httptest.NewRequest("PATCH", "/users/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(testing, http.StatusBadRequest, rr.Code)
}

func TestUpdateUser_Success(testing *testing.T) {
	repo, handler := setupUserTest()
	emptyContext := context.Background()

	_, err := repo.Create(emptyContext, &models.User{
		ID:         1,
		Username:   "oldname",
		Email:      "old@mail.com",
		FiscalCode: "RSSMRA80A01H501U",
	})
	if err != nil {
		return
	}

	router := chi.NewRouter()
	router.Patch("/users/{id}", handler.UpdateUser)

	body := `{"username": "newname", "email": "new@mail.com"}`
	req := httptest.NewRequest("PATCH", "/users/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(testing, http.StatusOK, rr.Code)

	var updated models.User
	err = json.NewDecoder(rr.Body).Decode(&updated)
	if err != nil {
		return
	}

	assert.Equal(testing, "newname", updated.Username)
	assert.Equal(testing, "new@mail.com", updated.Email)
}

func setupUserTest() (*FakeUserRepository, *UserHandler) {
	repository := NewFakeUserRepository()
	service := services.NewUserService(repository)
	handler := NewUserHandler(service)
	return repository, handler
}
