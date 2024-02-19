package users

type (
	CreateRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	UpdateRequest struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password"`
		IsActive bool   `json:"is_active"`
	}
)
