package services

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
	"user-service/internal/models"
	"user-service/internal/repositories"
)

type UserService struct {
	repository repositories.IUserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(repository repositories.IUserRepository) *UserService {
	return &UserService{repository: repository}
}

// CreateUser creates a new user
func (userService *UserService) CreateUser(context context.Context, user *models.User) (*models.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}
	if strings.TrimSpace(user.Username) == "" {
		return nil, errors.New("username is required")
	}
	if strings.TrimSpace(user.Password) == "" {
		return nil, errors.New("password is required")
	}
	if strings.TrimSpace(user.FiscalCode) == "" {
		return nil, errors.New("fiscal code is required")
	}

	user.Username = strings.ToLower(strings.TrimSpace(user.Username))
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))
	user.FiscalCode = strings.ToUpper(strings.TrimSpace(user.FiscalCode))

	if !isValidEmail(user.Email) {
		return nil, errors.New("invalid email")
	}

	var minimumLengthPassword = 10
	if len(user.Password) < minimumLengthPassword {
		return nil, errors.New("password must be at least 10 characters long")
	}

	if user.Role == "" {
		user.Role = "user"
	}

	if user.Role != "admin" && user.Role != "user" {
		return nil, errors.New("invalid role")
	}

	hashed, err := NewHashedPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(hashed)

	user.Username = strings.ToLower(strings.TrimSpace(user.Username))
	user.Password = hashed.Hash
	user.PasswordSalt = hashed.Salt

	return userService.repository.Create(context, user)
}

// isValidEmail validates an email address, ensuring it is non-empty, contains "@" and ".", and has proper structure.
func isValidEmail(email string) bool {
	email = strings.TrimSpace(strings.ToLower(email))
	if email == "" {
		return false
	}

	characterSeparator := "@"
	subParts := strings.Split(email, characterSeparator)

	numberOfSubParts := 2
	if len(subParts) != numberOfSubParts {
		return false
	}

	firstSubPartIndex := 0
	name := subParts[firstSubPartIndex]

	secondSubPartIndex := 1
	domain := subParts[secondSubPartIndex]

	if name == "" || domain == "" {
		return false
	}

	if !strings.Contains(domain, ".") {
		return false
	}

	return true
}

// GetAllUsers retrieves all users
func (userService *UserService) GetAllUsers(context context.Context) ([]*models.User, error) {
	if userService.repository == nil {
		return nil, errors.New("repository is nil")
	}
	return userService.repository.GetAll(context)
}

// DeleteUserByID deletes a user by their ID
func (userService *UserService) DeleteUserByID(context context.Context, id int64) error {
	if id <= 0 {
		return errors.New("id must be greater than 0")
	}
	return userService.repository.DeleteByID(context, id)
}

// GetUser retrieves a user by their ID
func (userService *UserService) GetUser(context context.Context, id int64) (*models.User, error) {
	return userService.repository.GetByID(context, id)
}

type HashedPassword struct {
	Hash string
	Salt string
}

// generateSalt generates a random salt of the specified number of bytes
// example generateSalt(16) => "a1b2c3d4e5f6g7h8"
func generateSalt(numberOfBytes int) (string, error) {
	bytes := make([]byte, numberOfBytes)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// hashPasswordWithSha256 hashes a given password using the SHA-256 algorithm and returns the hex-encoded string.
/*
	example: hashPasswordWithSha256("SuperSegreta123!", "a1b2c3d4e5f6g7h8") =>
	"9f3c8e0a7b1d2c4e5f6a7b8c9d0e1f2a3b4c5d67e8f90123456789abcdef0123"
*/
func hashPasswordWithSha256(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hash.Write([]byte(salt))
	return hex.EncodeToString(hash.Sum(nil))
}

// verifyPasswordSHA256 verifies if the provided password matches the expected hash with the given salt
func verifyPasswordSHA256(password string, salt string, expectedHash string) bool {
	computed := hashPasswordWithSha256(password, salt)
	return computed == expectedHash
}

// NewHashedPassword take a password and computes an hash with a specified salt
func NewHashedPassword(password string) (*HashedPassword, error) {
	var numberOfBytes = 16
	salt, err := generateSalt(numberOfBytes)
	if err != nil {
		return nil, err
	}

	hash := hashPasswordWithSha256(password, salt)

	return &HashedPassword{
		Hash: hash,
		Salt: salt,
	}, nil
}

// UpdateUser
func (userService *UserService) UpdateUser(context context.Context, id int64, request models.UpdateUserRequest) (*models.User, error) {
	user, err := userService.repository.GetByID(context, id)

	if err != nil {
		return nil, err
	}

	if request.Username != nil {
		user.Username = *request.Username
	}

	if request.Email != nil {
		if !isValidEmail(*request.Email) {
			return nil, errors.New("invalid email")
		}
		user.Email = *request.Email
	}

	if request.Telephone != nil {
		user.Telephone = *request.Telephone
	}

	if request.FiscalCode != nil {
		user.FiscalCode = *request.FiscalCode
	}

	if request.Role != nil {
		user.Role = *request.Role
	}
	var numberOfBytes = 16

	if request.Password != nil && *request.Password != "" {

		if looksLikeSHA256(*request.Password) {

			return nil, errors.New("password must be provided in plain text, not hashed")
		}
		salt, _ := generateSalt(numberOfBytes)
		hashed := hashPasswordWithSha256(*request.Password, salt)
		user.Password = hashed
		user.PasswordSalt = salt
	}

	if err := userService.repository.Update(context, user); err != nil {
		return nil, err
	}
	return user, nil

}

func looksLikeSHA256(password string) bool {
	if len(password) != 64 {
		return false
	}
	for _, c := range password {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return false
		}
	}
	return true
}
