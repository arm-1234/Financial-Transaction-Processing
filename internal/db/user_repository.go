package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"financial-transaction-system/internal/models"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type UserRepository struct {
	db *Database
}

func NewUserRepository(db *Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (id, email, password_hash, first_name, last_name, phone, date_of_birth, address, is_active, is_verified)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING created_at, updated_at`

	err := r.db.DB.QueryRow(
		query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.Phone,
		user.DateOfBirth,
		user.Address,
		user.IsActive,
		user.IsVerified,
	).Scan(&user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // unique_violation
				return fmt.Errorf("email already exists")
			}
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, password_hash, first_name, last_name, phone, date_of_birth, address, 
		       is_active, is_verified, created_at, updated_at
		FROM users 
		WHERE id = $1 AND is_active = true`

	err := r.db.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.DateOfBirth,
		&user.Address,
		&user.IsActive,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, password_hash, first_name, last_name, phone, date_of_birth, address, 
		       is_active, is_verified, created_at, updated_at
		FROM users 
		WHERE email = $1 AND is_active = true`

	err := r.db.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.DateOfBirth,
		&user.Address,
		&user.IsActive,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) Update(id uuid.UUID, req *models.UpdateUserRequest) (*models.User, error) {
	setParts := []string{}
	args := []interface{}{}
	argIndex := 1

	if req.FirstName != nil {
		setParts = append(setParts, fmt.Sprintf("first_name = $%d", argIndex))
		args = append(args, *req.FirstName)
		argIndex++
	}

	if req.LastName != nil {
		setParts = append(setParts, fmt.Sprintf("last_name = $%d", argIndex))
		args = append(args, *req.LastName)
		argIndex++
	}

	if req.Phone != nil {
		setParts = append(setParts, fmt.Sprintf("phone = $%d", argIndex))
		args = append(args, *req.Phone)
		argIndex++
	}

	if req.DateOfBirth != nil {
		setParts = append(setParts, fmt.Sprintf("date_of_birth = $%d", argIndex))
		args = append(args, *req.DateOfBirth)
		argIndex++
	}

	if req.Address != nil {
		setParts = append(setParts, fmt.Sprintf("address = $%d", argIndex))
		args = append(args, *req.Address)
		argIndex++
	}

	if len(setParts) == 0 {
		return r.GetByID(id)
	}

	setParts = append(setParts, fmt.Sprintf("updated_at = $%d", argIndex))
	args = append(args, time.Now())
	argIndex++

	query := fmt.Sprintf(`
		UPDATE users 
		SET %s
		WHERE id = $%d AND is_active = true
		RETURNING id, email, password_hash, first_name, last_name, phone, date_of_birth, address, 
		          is_active, is_verified, created_at, updated_at`,
		strings.Join(setParts, ", "),
		argIndex,
	)

	args = append(args, id)

	user := &models.User{}
	err := r.db.DB.QueryRow(query, args...).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.DateOfBirth,
		&user.Address,
		&user.IsActive,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) UpdatePassword(id uuid.UUID, passwordHash string) error {
	query := `UPDATE users SET password_hash = $1, updated_at = $2 WHERE id = $3 AND is_active = true`

	result, err := r.db.DB.Exec(query, passwordHash, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (r *UserRepository) Deactivate(id uuid.UUID) error {
	query := `UPDATE users SET is_active = false, updated_at = $1 WHERE id = $2`

	result, err := r.db.DB.Exec(query, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to deactivate user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (r *UserRepository) SetVerified(id uuid.UUID, verified bool) error {
	query := `UPDATE users SET is_verified = $1, updated_at = $2 WHERE id = $3 AND is_active = true`

	result, err := r.db.DB.Exec(query, verified, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to update verification status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
