package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"project-dpwims/shared/utilities"
	"strconv"
	"users-service/pkg/models"
	"users-service/pkg/repositories"
	"users-service/pkg/services"

	"github.com/go-chi/chi/v5"
)

const KeyContentType = "Content-Type"
const ValueAppJson = "application/json"
const baseNumber = 10
const bitSize = 64
const errorMessageUserNotFound = "user not found"
const errorMessageInvalidID = "invalid id"

// UserHandler struct to handle the user related HTTP requests
type UserHandler struct {
	service *services.UserService
}

// NewUserHandler to create an instance of UserHandler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{service: userService}
}

// CreateUser a handlers method to create a new user into repositories memory
func (userHandler *UserHandler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	decodeJson, ok := utilities.DecodeJSON[models.CreateUserRequest](writer, request)
	if !ok {
		return
	}

	user := &models.User{
		Username:   decodeJson.Username,
		Password:   decodeJson.Password,
		Email:      decodeJson.Email,
		FiscalCode: decodeJson.FiscalCode,
		Telephone:  decodeJson.Telephone,
		Role:       decodeJson.Role,
	}

	created, err := userHandler.service.CreateUser(request.Context(), user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	utilities.WriteJSON(writer, http.StatusCreated, created)
}

// GetUser a handlers method to get a user by id from repositories memory
func (userHandler *UserHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	id, ok := utilities.ParseIDParam(writer, request, "id")
	if !ok {
		return
	}
	user, err := userHandler.service.GetUser(request.Context(), id)
	if err != nil || user == nil {
		http.Error(writer, errorMessageUserNotFound, http.StatusNotFound)
		return
	}
	utilities.WriteJSON(writer, http.StatusOK, user)
}

// GetAllUsers a handlers method to get all users into repositories memory
func (userHandler *UserHandler) GetAllUsers(writer http.ResponseWriter, request *http.Request) {
	users, err := userHandler.service.GetAllUsers(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(users)
}

// DeleteUser a handlers method to delete a user by id from repositories memory
func (userHandler *UserHandler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idString, baseNumber, bitSize)
	if err != nil {
		http.Error(writer, errorMessageInvalidID, http.StatusBadRequest)
		return
	}
	err = userHandler.service.DeleteUserByID(request.Context(), id)
	if err != nil {
		http.Error(writer, errorMessageUserNotFound, http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

// UpdateUser a handlers method to update a user by id from repositories memory
func (userHandler *UserHandler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idString, baseNumber, bitSize)
	if err != nil || id <= 0 {
		http.Error(writer, errorMessageInvalidID, http.StatusBadRequest)
		return
	}

	var updateUserRequest models.UpdateUserRequest
	if err := json.NewDecoder(request.Body).Decode(&updateUserRequest); err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}

	updatedUser, err := userHandler.service.UpdateUser(request.Context(), id, updateUserRequest)
	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			http.Error(writer, errorMessageUserNotFound, http.StatusNotFound)
			return
		}

		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(updatedUser)
	if err != nil {
		return
	}
}
