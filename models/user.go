package models

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"-" validate:"required,min=6"`
	Avatar   string `json:"avatar"`
}
