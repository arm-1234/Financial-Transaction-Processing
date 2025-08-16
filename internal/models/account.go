package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccountType string
type AccountStatus string

const (
	AccountTypeSavings    AccountType = "savings"
	AccountTypeChecking   AccountType = "checking"
	AccountTypeBusiness   AccountType = "business"
	AccountTypeInvestment AccountType = "investment"
)

const (
	AccountStatusActive    AccountStatus = "active"
	AccountStatusSuspended AccountStatus = "suspended"
	AccountStatusClosed    AccountStatus = "closed"
)

type Account struct {
	ID               uuid.UUID       `json:"id" db:"id"`
	UserID           uuid.UUID       `json:"user_id" db:"user_id"`
	AccountNumber    string          `json:"account_number" db:"account_number"`
	AccountType      AccountType     `json:"account_type" db:"account_type"`
	AccountName      string          `json:"account_name" db:"account_name"`
	Balance          decimal.Decimal `json:"balance" db:"balance"`
	AvailableBalance decimal.Decimal `json:"available_balance" db:"available_balance"`
	Currency         string          `json:"currency" db:"currency"`
	Status           AccountStatus   `json:"status" db:"status"`
	DailyLimit       decimal.Decimal `json:"daily_limit" db:"daily_limit"`
	MonthlyLimit     decimal.Decimal `json:"monthly_limit" db:"monthly_limit"`
	IsPrimary        bool            `json:"is_primary" db:"is_primary"`
	CreatedAt        time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" db:"updated_at"`
}

type CreateAccountRequest struct {
	AccountType  AccountType     `json:"account_type" validate:"required,oneof=savings checking business investment"`
	AccountName  string          `json:"account_name" validate:"required,min=3,max=100"`
	Currency     string          `json:"currency" validate:"required,len=3"`
	DailyLimit   decimal.Decimal `json:"daily_limit" validate:"required,gt=0"`
	MonthlyLimit decimal.Decimal `json:"monthly_limit" validate:"required,gt=0"`
	IsPrimary    bool            `json:"is_primary"`
}

type UpdateAccountRequest struct {
	AccountName  *string          `json:"account_name,omitempty" validate:"omitempty,min=3,max=100"`
	Status       *AccountStatus   `json:"status,omitempty" validate:"omitempty,oneof=active suspended closed"`
	DailyLimit   *decimal.Decimal `json:"daily_limit,omitempty" validate:"omitempty,gt=0"`
	MonthlyLimit *decimal.Decimal `json:"monthly_limit,omitempty" validate:"omitempty,gt=0"`
	IsPrimary    *bool            `json:"is_primary,omitempty"`
}

type AccountBalance struct {
	AccountID        uuid.UUID       `json:"account_id"`
	AccountNumber    string          `json:"account_number"`
	Balance          decimal.Decimal `json:"balance"`
	AvailableBalance decimal.Decimal `json:"available_balance"`
	Currency         string          `json:"currency"`
	LastUpdated      time.Time       `json:"last_updated"`
}

type AccountSummary struct {
	ID               uuid.UUID       `json:"id"`
	AccountNumber    string          `json:"account_number"`
	AccountType      AccountType     `json:"account_type"`
	AccountName      string          `json:"account_name"`
	Balance          decimal.Decimal `json:"balance"`
	AvailableBalance decimal.Decimal `json:"available_balance"`
	Currency         string          `json:"currency"`
	Status           AccountStatus   `json:"status"`
	IsPrimary        bool            `json:"is_primary"`
}

// AccountWithUser includes user information for admin views
type AccountWithUser struct {
	Account
	User UserProfile `json:"user"`
}
