package auth

import (
	"learn-go-fiber/modules/user"
	"time"
)

// RegisterResponse ...
type RegisterResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoginResponse ...
type LoginResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FormatRegisterResponse ...
func FormatRegisterResponse(user *user.User) *RegisterResponse {
	formatter := &RegisterResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return formatter
}

// FormatLoginResponse ...
func FormatLoginResponse(user *user.User, token string) *LoginResponse {
	formatter := &LoginResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return formatter
}
