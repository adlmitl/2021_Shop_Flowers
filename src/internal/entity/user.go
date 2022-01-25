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
	Id       uuid.UUID `json:"id" db:"id"`
	Login    string    `json:"login" db:"login"`
	Password string    `json:"password" db:"password"`
}
