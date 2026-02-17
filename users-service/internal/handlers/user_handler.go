package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"user-service/internal/models"
	"user-service/internal/repositories"
	"user-service/internal/services"

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

// CreateUser a handler method to create a new user into repository memory
func (userHandler *UserHandler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "invalid JSON body"+err.Error(), http.StatusBadRequest)
		return
	}

	created, err := userHandler.service.CreateUser(request.Context(), &user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(created)
}

// GetUser a handler method to get a user by id from repository memory
func (userHandler *UserHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idStr, baseNumber, bitSize)
	user, err := userHandler.service.GetUser(request.Context(), id)

	if err != nil || user == nil {
		http.Error(writer, errorMessageUserNotFound, http.StatusNotFound)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(user)
}

// GetAllUsers a handler method to get all users into repository memory
func (userHandler *UserHandler) GetAllUsers(writer http.ResponseWriter, request *http.Request) {
	users, err := userHandler.service.GetAllUsers(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(users)
}

// DeleteUser a handler method to delete a user by id from repository memory
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

// UpdateUser a handler method to update a user by id from repository memory
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
