package auth

import (
	"errors"
	"time"

	"financial-transaction-system/internal/config"
	"financial-transaction-system/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrExpiredToken  = errors.New("token has expired")
	ErrInvalidClaims = errors.New("invalid token claims")
)

type JWTManager struct {
	secretKey     string
	tokenExpiry   time.Duration
	refreshExpiry time.Duration
}

type Claims struct {
	UserID    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	TokenType string    `json:"token_type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

func NewJWTManager(cfg *config.Config) *JWTManager {
	return &JWTManager{
		secretKey:     cfg.JWT.Secret,
		tokenExpiry:   cfg.GetJWTExpiry(),
		refreshExpiry: cfg.GetJWTRefreshExpiry(),
	}
}

func (j *JWTManager) GenerateTokens(user *models.User) (*models.LoginResponse, error) {
	accessToken, err := j.generateToken(user, "access", j.tokenExpiry)
	if err != nil {
		return nil, err
	}

	refreshToken, err := j.generateToken(user, "refresh", j.refreshExpiry)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(j.tokenExpiry).Unix(),
	}, nil
}

func (j *JWTManager) generateToken(user *models.User, tokenType string, expiry time.Duration) (string, error) {
	now := time.Now()
	claims := &Claims{
		UserID:    user.ID,
		Email:     user.Email,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "financial-transaction-system",
			Subject:   user.ID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTManager) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidClaims
	}

	return claims, nil
}

func (j *JWTManager) RefreshToken(refreshTokenString string) (*models.LoginResponse, error) {
	claims, err := j.ValidateToken(refreshTokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != "refresh" {
		return nil, ErrInvalidToken
	}

	user := &models.User{
		ID:    claims.UserID,
		Email: claims.Email,
	}

	return j.GenerateTokens(user)
}

func (j *JWTManager) ExtractUserID(tokenString string) (uuid.UUID, error) {
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		return uuid.Nil, err
	}

	if claims.TokenType != "access" {
		return uuid.Nil, ErrInvalidToken
	}

	return claims.UserID, nil
}
