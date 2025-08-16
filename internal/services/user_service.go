package services

import (
	"fmt"

	"financial-transaction-system/internal/auth"
	"financial-transaction-system/internal/db"
	"financial-transaction-system/internal/models"
	"financial-transaction-system/internal/utils"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo   *db.UserRepository
	jwtManager *auth.JWTManager
}

func NewUserService(userRepo *db.UserRepository, jwtManager *auth.JWTManager) *UserService {
	return &UserService{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

func (s *UserService) Register(req *models.CreateUserRequest) (*models.LoginResponse, error) {
	// Validate password
	if !utils.IsValidPassword(req.Password) {
		return nil, fmt.Errorf("password must be at least %d characters long", utils.MinPasswordLength)
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user object
	user := &models.User{
		ID:           uuid.New(),
		Email:        req.Email,
		PasswordHash: hashedPassword,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Phone:        req.Phone,
		DateOfBirth:  req.DateOfBirth,
		Address:      req.Address,
		IsActive:     true,
		IsVerified:   false,
	}

	// Save user to database
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Generate JWT tokens
	return s.jwtManager.GenerateTokens(user)
}

func (s *UserService) Login(req *models.LoginRequest) (*models.LoginResponse, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	// Verify password
	if err := utils.VerifyPassword(user.PasswordHash, req.Password); err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, fmt.Errorf("account is deactivated")
	}

	// Generate JWT tokens
	return s.jwtManager.GenerateTokens(user)
}

func (s *UserService) RefreshToken(req *models.RefreshTokenRequest) (*models.LoginResponse, error) {
	return s.jwtManager.RefreshToken(req.RefreshToken)
}

func (s *UserService) GetProfile(userID uuid.UUID) (*models.UserProfile, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return &models.UserProfile{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *UserService) UpdateProfile(userID uuid.UUID, req *models.UpdateUserRequest) (*models.UserProfile, error) {
	user, err := s.userRepo.Update(userID, req)
	if err != nil {
		return nil, err
	}

	return &models.UserProfile{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *UserService) ChangePassword(userID uuid.UUID, currentPassword, newPassword string) error {
	// Get user
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	// Verify current password
	if err := utils.VerifyPassword(user.PasswordHash, currentPassword); err != nil {
		return fmt.Errorf("current password is incorrect")
	}

	// Validate new password
	if !utils.IsValidPassword(newPassword) {
		return fmt.Errorf("new password must be at least %d characters long", utils.MinPasswordLength)
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash new password: %w", err)
	}

	// Update password
	return s.userRepo.UpdatePassword(userID, hashedPassword)
}

func (s *UserService) DeactivateAccount(userID uuid.UUID) error {
	return s.userRepo.Deactivate(userID)
}

func (s *UserService) VerifyAccount(userID uuid.UUID) error {
	return s.userRepo.SetVerified(userID, true)
}
