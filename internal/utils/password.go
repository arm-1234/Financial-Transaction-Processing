package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	MinPasswordLength = 8  // Minimum password length
	DefaultCost       = 12 // bcrypt cost factor
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// VerifyPassword compares hashed and plain passwords
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// IsValidPassword checks minimum password requirements
func IsValidPassword(password string) bool {
	return len(password) >= MinPasswordLength
}
