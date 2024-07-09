package models

type User struct {
	Id       string `json:"pid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
