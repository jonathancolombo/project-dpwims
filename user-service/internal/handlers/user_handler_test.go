package handlers

import (
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

func setupUserTest() (*repositories.InMemoryUserRepository, *UserHandler) {
	repository := repositories.NewInMemoryUserRepository()
	service := services.NewUserService(repository)
	handler := NewUserHandler(service)
	return repository, handler
}
