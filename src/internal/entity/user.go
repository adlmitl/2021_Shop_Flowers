package entity

import (
	"github.com/google/uuid"
)

// UsersList - List users.
type UsersList struct {
	Users []*User
}

// User - Struct.
type User struct {
	Id       uuid.UUID `json:"id"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}
