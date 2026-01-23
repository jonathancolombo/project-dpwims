package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-service/internal/models"
	"user-service/internal/services"

	"github.com/go-chi/chi/v5"
)

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

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(created)
	if err != nil {
		return
	}
}

func (userHandler *UserHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	user, err := userHandler.service.GetUser(id)

	if err != nil || user == nil {
		http.Error(writer, "user not found", http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(user)
	checkStatusInternalServerError(writer, err)
}

// GetAllUsers a handler method to get all users into repository memory
func (userHandler *UserHandler) GetAllUsers(writer http.ResponseWriter, _ *http.Request) {
	users, err := userHandler.service.GetAllUsers()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(users)
	checkStatusInternalServerError(writer, err)
}

func (userHandler *UserHandler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	idStr := request.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
		return
	}
	err = userHandler.service.DeleteUserByID(id)
	if err != nil {
		http.Error(writer, "user not found", http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent) // 204
}

func checkStatusInternalServerError(writer http.ResponseWriter, err error) {
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
