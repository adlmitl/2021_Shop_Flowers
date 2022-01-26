package entity

import (
	"github.com/google/uuid"
)

// UsersList - список пользователей.
type UsersList struct {
	Users []*User
}

// User - сущность.
type User struct {
	Id       uuid.UUID `json:"id"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}
