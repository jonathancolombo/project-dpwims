package models

// UpdateUserRequest a dto model to updating fields
type UpdateUserRequest struct {
	Username   *string `json:"username"`
	Email      *string `json:"email"`
	Telephone  *string `json:"telephone"`
	FiscalCode *string `json:"fiscal_code"`
	Role       *string `json:"role"`
	Password   *string `json:"password"`
}
