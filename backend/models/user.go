package models

type User struct {
	UUID     string `json:"uuid" db:"uuid"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}
