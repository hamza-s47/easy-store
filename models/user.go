package models

import "time"

type User struct {
	ID        int       `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	DOB       time.Time `db:"dob"`
	Phone     *string   `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
}

type UserRegister struct {
	Email           string  `json:"email"`
	Password        string  `json:"password"`
	ConfirmPassword string  `json:"confirm_password"`
	DOB             string  `json:"dob"`
	Phone           *string `json:"phone"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
