package models

type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	FiscalCode   string `json:"fiscal_code"`
	Telephone    string `json:"telephone"`
	Role         string `json:"role"`
	PasswordSalt string `json:"-"`
}
