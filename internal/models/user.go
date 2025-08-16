package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	Email        string     `json:"email" db:"email" validate:"required,email"`
	PasswordHash string     `json:"-" db:"password_hash"`
	FirstName    string     `json:"first_name" db:"first_name" validate:"required,min=2,max=50"`
	LastName     string     `json:"last_name" db:"last_name" validate:"required,min=2,max=50"`
	Phone        *string    `json:"phone,omitempty" db:"phone" validate:"omitempty,min=10,max=20"`
	DateOfBirth  *time.Time `json:"date_of_birth,omitempty" db:"date_of_birth"`
	Address      *string    `json:"address,omitempty" db:"address"`
	IsActive     bool       `json:"is_active" db:"is_active"`
	IsVerified   bool       `json:"is_verified" db:"is_verified"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}

type CreateUserRequest struct {
	Email       string     `json:"email" validate:"required,email"`
	Password    string     `json:"password" validate:"required,min=8"`
	FirstName   string     `json:"first_name" validate:"required,min=2,max=50"`
	LastName    string     `json:"last_name" validate:"required,min=2,max=50"`
	Phone       *string    `json:"phone,omitempty" validate:"omitempty,min=10,max=20"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	Address     *string    `json:"address,omitempty"`
}

type UpdateUserRequest struct {
	FirstName   *string    `json:"first_name,omitempty" validate:"omitempty,min=2,max=50"`
	LastName    *string    `json:"last_name,omitempty" validate:"omitempty,min=2,max=50"`
	Phone       *string    `json:"phone,omitempty" validate:"omitempty,min=10,max=20"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	Address     *string    `json:"address,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User         User   `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// UserProfile represents public user information
type UserProfile struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     *string   `json:"phone,omitempty"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
