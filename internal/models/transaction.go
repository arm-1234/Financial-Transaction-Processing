package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionType string
type TransactionStatus string

const (
	TransactionTypeTransfer   TransactionType = "transfer"
	TransactionTypeDeposit    TransactionType = "deposit"
	TransactionTypeWithdrawal TransactionType = "withdrawal"
	TransactionTypeFee        TransactionType = "fee"
	TransactionTypeInterest   TransactionType = "interest"
	TransactionTypeRefund     TransactionType = "refund"
)

const (
	TransactionStatusPending    TransactionStatus = "pending"
	TransactionStatusProcessing TransactionStatus = "processing"
	TransactionStatusCompleted  TransactionStatus = "completed"
	TransactionStatusFailed     TransactionStatus = "failed"
	TransactionStatusCancelled  TransactionStatus = "cancelled"
	TransactionStatusReversed   TransactionStatus = "reversed"
)

type Transaction struct {
	ID                uuid.UUID         `json:"id" db:"id"`
	TransactionNumber string            `json:"transaction_number" db:"transaction_number"`
	FromAccountID     *uuid.UUID        `json:"from_account_id,omitempty" db:"from_account_id"`
	ToAccountID       *uuid.UUID        `json:"to_account_id,omitempty" db:"to_account_id"`
	TransactionType   TransactionType   `json:"transaction_type" db:"transaction_type"`
	Amount            decimal.Decimal   `json:"amount" db:"amount"`
	Currency          string            `json:"currency" db:"currency"`
	ExchangeRate      decimal.Decimal   `json:"exchange_rate" db:"exchange_rate"`
	Fee               decimal.Decimal   `json:"fee" db:"fee"`
	Description       *string           `json:"description,omitempty" db:"description"`
	ReferenceNumber   *string           `json:"reference_number,omitempty" db:"reference_number"`
	Status            TransactionStatus `json:"status" db:"status"`
	ProcessedAt       *time.Time        `json:"processed_at,omitempty" db:"processed_at"`
	CreatedAt         time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at" db:"updated_at"`
}

type TransferRequest struct {
	FromAccountID   uuid.UUID       `json:"from_account_id" validate:"required"`
	ToAccountID     uuid.UUID       `json:"to_account_id" validate:"required"`
	Amount          decimal.Decimal `json:"amount" validate:"required,gt=0"`
	Currency        string          `json:"currency" validate:"required,len=3"`
	Description     *string         `json:"description,omitempty" validate:"omitempty,max=500"`
	ReferenceNumber *string         `json:"reference_number,omitempty" validate:"omitempty,max=50"`
}

type DepositRequest struct {
	ToAccountID     uuid.UUID       `json:"to_account_id" validate:"required"`
	Amount          decimal.Decimal `json:"amount" validate:"required,gt=0"`
	Currency        string          `json:"currency" validate:"required,len=3"`
	Description     *string         `json:"description,omitempty" validate:"omitempty,max=500"`
	ReferenceNumber *string         `json:"reference_number,omitempty" validate:"omitempty,max=50"`
}

type WithdrawalRequest struct {
	FromAccountID   uuid.UUID       `json:"from_account_id" validate:"required"`
	Amount          decimal.Decimal `json:"amount" validate:"required,gt=0"`
	Currency        string          `json:"currency" validate:"required,len=3"`
	Description     *string         `json:"description,omitempty" validate:"omitempty,max=500"`
	ReferenceNumber *string         `json:"reference_number,omitempty" validate:"omitempty,max=50"`
}

type TransactionResponse struct {
	ID                uuid.UUID         `json:"id"`
	TransactionNumber string            `json:"transaction_number"`
	TransactionType   TransactionType   `json:"transaction_type"`
	Amount            decimal.Decimal   `json:"amount"`
	Currency          string            `json:"currency"`
	Fee               decimal.Decimal   `json:"fee"`
	Description       *string           `json:"description,omitempty"`
	Status            TransactionStatus `json:"status"`
	ProcessedAt       *time.Time        `json:"processed_at,omitempty"`
	CreatedAt         time.Time         `json:"created_at"`
	FromAccount       *AccountSummary   `json:"from_account,omitempty"`
	ToAccount         *AccountSummary   `json:"to_account,omitempty"`
}

type TransactionHistory struct {
	Transactions []TransactionResponse `json:"transactions"`
	Pagination   PaginationResponse    `json:"pagination"`
}

type TransactionStatement struct {
	AccountID      uuid.UUID             `json:"account_id"`
	AccountNumber  string                `json:"account_number"`
	StartDate      time.Time             `json:"start_date"`
	EndDate        time.Time             `json:"end_date"`
	OpeningBalance decimal.Decimal       `json:"opening_balance"`
	ClosingBalance decimal.Decimal       `json:"closing_balance"`
	Transactions   []TransactionResponse `json:"transactions"`
	Summary        TransactionSummary    `json:"summary"`
}

type TransactionSummary struct {
	TotalCredits     decimal.Decimal `json:"total_credits"`
	TotalDebits      decimal.Decimal `json:"total_debits"`
	TotalFees        decimal.Decimal `json:"total_fees"`
	TransactionCount int             `json:"transaction_count"`
}

type PaginationRequest struct {
	Page     int `json:"page" query:"page" validate:"min=1"`
	PageSize int `json:"page_size" query:"page_size" validate:"min=1,max=100"`
}

type PaginationResponse struct {
	Page         int   `json:"page"`
	PageSize     int   `json:"page_size"`
	TotalPages   int   `json:"total_pages"`
	TotalRecords int64 `json:"total_records"`
	HasNext      bool  `json:"has_next"`
	HasPrevious  bool  `json:"has_previous"`
}

// TransactionFilter for filtering transaction history
type TransactionFilter struct {
	AccountID       *uuid.UUID         `json:"account_id,omitempty" query:"account_id"`
	TransactionType *TransactionType   `json:"transaction_type,omitempty" query:"transaction_type"`
	Status          *TransactionStatus `json:"status,omitempty" query:"status"`
	StartDate       *time.Time         `json:"start_date,omitempty" query:"start_date"`
	EndDate         *time.Time         `json:"end_date,omitempty" query:"end_date"`
	MinAmount       *decimal.Decimal   `json:"min_amount,omitempty" query:"min_amount"`
	MaxAmount       *decimal.Decimal   `json:"max_amount,omitempty" query:"max_amount"`
	Currency        *string            `json:"currency,omitempty" query:"currency"`
	PaginationRequest
}
