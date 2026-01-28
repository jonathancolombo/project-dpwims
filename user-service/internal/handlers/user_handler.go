package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-service/internal/models"
	"user-service/internal/services"

	"github.com/go-chi/chi/v5"
)

const KeyContentType = "Content-Type"
const ValueAppJson = "application/json"
const baseNumber = 10
const bitSize = 64

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{service: userService}
}

func (userHandler *UserHandler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "invalid JSON body"+err.Error(), http.StatusBadRequest)
		return
	}

	created, err := userHandler.service.CreateUser(&user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(created)
}

func (userHandler *UserHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idStr, baseNumber, bitSize)
	user, err := userHandler.service.GetUser(id)

	if err != nil || user == nil {
		http.Error(writer, "user not found", http.StatusNotFound)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(user)
}

// GetAllUsers a handler method to get all users into repository memory
func (userHandler *UserHandler) GetAllUsers(writer http.ResponseWriter, _ *http.Request) {
	users, err := userHandler.service.GetAllUsers()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(users)
}

func (userHandler *UserHandler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idString, baseNumber, bitSize)
	if err != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
		return
	}
	err = userHandler.service.DeleteUserByID(id)
	if err != nil {
		http.Error(writer, "user not found", http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}
