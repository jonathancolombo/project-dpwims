package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"users-service/pkg/models"
)

var ErrUserNotFound = errors.New("user not found")

// MySQLUserRepository provides methods for CRUD operations on the 'users' table in a MySQL database.
type MySQLUserRepository struct {
	database *sql.DB
}

// NewMySQLRepositoryUsers initializes a new MySQLUserRepository with the provided database connection.
func NewMySQLRepositoryUsers(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{database: db}
}

// Create a method to create a user and save into a db
func (mySqlUserRepository *MySQLUserRepository) Create(context context.Context, user *models.User) (*models.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	query := `INSERT INTO users (username, password, password_salt, email, fiscal_code, telephone, role) VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := mySqlUserRepository.database.PrepareContext(context, query)

	if err != nil {
		_ = fmt.Errorf("prepare statement: %w", err)
		return nil, err
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			return
		}
	}(statement)

	result, err := statement.Exec(
		strings.ToLower(user.Username),
		user.Password,
		user.PasswordSalt,
		user.Email,
		user.FiscalCode,
		user.Telephone,
		user.Role,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to insert id: %w", err)
	}
	user.ID = id
	return user, nil
}

// GetByID is a method to find the right user using id field
func (mySqlUserRepository *MySQLUserRepository) GetByID(context context.Context, id int64) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password, email, fiscal_code, telephone, role FROM users WHERE id = ?`
	row := mySqlUserRepository.database.QueryRowContext(context, query, id)

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FiscalCode, &user.Telephone, &user.Role)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("scan user: %w", err)
	}

	return &user, nil
}

// GetAll retrieves all users into a slice
func (mySqlUserRepository *MySQLUserRepository) GetAll(context context.Context) ([]*models.User, error) {
	var users []*models.User

	query := `SELECT id, username, password, email, fiscal_code, telephone, role FROM users`

	rows, err := mySqlUserRepository.database.QueryContext(context, query)
	if err != nil {
		return nil, fmt.Errorf("query all users: %w", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FiscalCode, &user.Telephone, &user.Role); err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}
		users = append(users, &user)
	}

	return users, nil
}

// DeleteByID delete a user by his id
func (mySqlUserRepository *MySQLUserRepository) DeleteByID(context context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := mySqlUserRepository.database.ExecContext(context, query, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rows == 0 {
		return ErrUserNotFound
	}

	return nil
}

// Update a method that's update a current user into db
func (mySqlUserRepository *MySQLUserRepository) Update(context context.Context, user *models.User) error {
	query := `
        UPDATE users
        SET username = ?, email = ?, telephone = ?, fiscal_code = ?, role = ?, password = ?, password_salt = ?
        WHERE id = ?
    `
	_, err := mySqlUserRepository.database.ExecContext(context, query,
		user.Username,
		user.Email,
		user.Telephone,
		user.FiscalCode,
		user.Role,
		user.Password,
		user.PasswordSalt,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	return nil
}

// GetByEmail is a method that's retrieve a user by his email
func (mySqlUserRepository *MySQLUserRepository) GetByEmail(context context.Context, email string) (*models.User, error) {
	var user models.User

	query := `
        SELECT id, username, email, password, password_salt, role, fiscal_code, telephone
        FROM users
        WHERE email = ?
        LIMIT 1
    `

	row := mySqlUserRepository.database.QueryRowContext(context, query, email)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.PasswordSalt,
		&user.Role,
		&user.FiscalCode,
		&user.Telephone,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
