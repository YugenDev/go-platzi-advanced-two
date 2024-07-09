package models

type User struct {
	Id       int64  `json:"pid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
