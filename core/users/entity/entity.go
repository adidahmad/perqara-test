package entity

import (
	"time"
)

type (
	Users struct {
		ID        uint       `json:"id"`
		Email     string     `json:"email"`
		Password  string     `json:"password"`
		IsActive  bool       `json:"is_active"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}
)

func (Users) TableName() string {
	return "users"
}
