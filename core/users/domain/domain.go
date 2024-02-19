package domain

import "time"

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

type (
	CreateUserRequest struct {
		Email    string
		Password string
		IsActive bool
	}

	UpdateUserRequest struct {
		Email    string
		Password string
		IsActive bool
	}
)
