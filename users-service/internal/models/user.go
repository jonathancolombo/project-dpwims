package models

type User struct {
	ID           int64    `json:"id"`
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	Email        string   `json:"email"`
	FiscalCode   string   `json:"fiscal_code"`
	Telephone    string   `json:"telephone"`
	Role         UserRole `json:"role"`
	PasswordSalt string   `json:"-"`
}

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleCustomer UserRole = "customer"
)

// UpdateUserRequest a dto model to updating fields
type UpdateUserRequest struct {
	Username   *string   `json:"username"`
	Email      *string   `json:"email"`
	Telephone  *string   `json:"telephone"`
	FiscalCode *string   `json:"fiscal_code"`
	Role       *UserRole `json:"role"`
	Password   *string   `json:"password"`
}
