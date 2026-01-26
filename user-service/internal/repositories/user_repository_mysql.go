package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"user-service/internal/models"
)

type MySQLUserRepository struct {
	database *sql.DB
}

func NewMySQLRepositoryUsers(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{database: db}
}

func (mySqlUserRepository *MySQLUserRepository) Create(context context.Context, user *models.User) (*models.User, error) {
	query := `INSERT INTO users 
    			(username, password, email, fiscal_code, telephone) 
					VALUES (?, ?, ?, ?, ?)`
	statement, err := mySqlUserRepository.database.PrepareContext(context, query)

	if err != nil {
		return nil, err
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			return
		}
	}(statement)

	result, err := statement.Exec(user.Username, user.Password, user.Email, user.Telephone)
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

/*
func (r *MySQLUserRepository) FindByID(ctx context.Context, id int64) (*models.User, error) {
	query := `
        SELECT id, username, password, email, fiscal_code, telephone
        FROM users
        WHERE id = ?
    `

	row := r.database.QueryRowContext(ctx, query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FiscalCode, &user.Telephone)
	if err == sql.ErrNoRows {
		return nil, repositories.ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan user: %w", err)
	}

	return &user, nil
}

func (r *MySQLUserRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	query := `
        SELECT id, username, password, email, fiscal_code, telephone
        FROM users
    `

	rows, err := r.database.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query all users: %w", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FiscalCode, &user.Telephone); err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *MySQLUserRepository) DeleteByID(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := r.database.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rows == 0 {
		return repositories.ErrUserNotFound
	}

	return nil
}
*/
