package models

import (
	"encoding/json"
	"net"
	"time"

	"github.com/google/uuid"
)

type AuditAction string

const (
	AuditActionUserCreated          AuditAction = "user_created"
	AuditActionUserUpdated          AuditAction = "user_updated"
	AuditActionUserDeleted          AuditAction = "user_deleted"
	AuditActionUserLogin            AuditAction = "user_login"
	AuditActionUserLogout           AuditAction = "user_logout"
	AuditActionAccountCreated       AuditAction = "account_created"
	AuditActionAccountUpdated       AuditAction = "account_updated"
	AuditActionAccountSuspended     AuditAction = "account_suspended"
	AuditActionAccountClosed        AuditAction = "account_closed"
	AuditActionTransactionCreated   AuditAction = "transaction_created"
	AuditActionTransactionProcessed AuditAction = "transaction_processed"
	AuditActionTransactionFailed    AuditAction = "transaction_failed"
	AuditActionTransactionCancelled AuditAction = "transaction_cancelled"
	AuditActionBalanceUpdated       AuditAction = "balance_updated"
	AuditActionFraudDetected        AuditAction = "fraud_detected"
	AuditActionFraudCleared         AuditAction = "fraud_cleared"
	AuditActionLimitExceeded        AuditAction = "limit_exceeded"
	AuditActionPasswordChanged      AuditAction = "password_changed"
	AuditActionEmailChanged         AuditAction = "email_changed"
	AuditActionProfileUpdated       AuditAction = "profile_updated"
)

type AuditLog struct {
	ID            uuid.UUID       `json:"id" db:"id"`
	UserID        *uuid.UUID      `json:"user_id,omitempty" db:"user_id"`
	AccountID     *uuid.UUID      `json:"account_id,omitempty" db:"account_id"`
	TransactionID *uuid.UUID      `json:"transaction_id,omitempty" db:"transaction_id"`
	Action        AuditAction     `json:"action" db:"action"`
	EntityType    string          `json:"entity_type" db:"entity_type"`
	EntityID      *uuid.UUID      `json:"entity_id,omitempty" db:"entity_id"`
	OldValues     json.RawMessage `json:"old_values,omitempty" db:"old_values"`
	NewValues     json.RawMessage `json:"new_values,omitempty" db:"new_values"`
	IPAddress     *net.IP         `json:"ip_address,omitempty" db:"ip_address"`
	UserAgent     *string         `json:"user_agent,omitempty" db:"user_agent"`
	SessionID     *string         `json:"session_id,omitempty" db:"session_id"`
	Description   *string         `json:"description,omitempty" db:"description"`
	Metadata      json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
}

type CreateAuditLogRequest struct {
	UserID        *uuid.UUID  `json:"user_id,omitempty"`
	AccountID     *uuid.UUID  `json:"account_id,omitempty"`
	TransactionID *uuid.UUID  `json:"transaction_id,omitempty"`
	Action        AuditAction `json:"action"`
	EntityType    string      `json:"entity_type"`
	EntityID      *uuid.UUID  `json:"entity_id,omitempty"`
	OldValues     interface{} `json:"old_values,omitempty"`
	NewValues     interface{} `json:"new_values,omitempty"`
	IPAddress     *net.IP     `json:"ip_address,omitempty"`
	UserAgent     *string     `json:"user_agent,omitempty"`
	SessionID     *string     `json:"session_id,omitempty"`
	Description   *string     `json:"description,omitempty"`
	Metadata      interface{} `json:"metadata,omitempty"`
}

type AuditLogFilter struct {
	UserID        *uuid.UUID   `json:"user_id,omitempty" query:"user_id"`
	AccountID     *uuid.UUID   `json:"account_id,omitempty" query:"account_id"`
	TransactionID *uuid.UUID   `json:"transaction_id,omitempty" query:"transaction_id"`
	Action        *AuditAction `json:"action,omitempty" query:"action"`
	EntityType    *string      `json:"entity_type,omitempty" query:"entity_type"`
	EntityID      *uuid.UUID   `json:"entity_id,omitempty" query:"entity_id"`
	StartDate     *time.Time   `json:"start_date,omitempty" query:"start_date"`
	EndDate       *time.Time   `json:"end_date,omitempty" query:"end_date"`
	IPAddress     *net.IP      `json:"ip_address,omitempty" query:"ip_address"`
	PaginationRequest
}

// Fraud Alert Models
type FraudSeverity string
type FraudStatus string

const (
	FraudSeverityLow      FraudSeverity = "low"
	FraudSeverityMedium   FraudSeverity = "medium"
	FraudSeverityHigh     FraudSeverity = "high"
	FraudSeverityCritical FraudSeverity = "critical"
)

const (
	FraudStatusOpen          FraudStatus = "open"
	FraudStatusInvestigating FraudStatus = "investigating"
	FraudStatusResolved      FraudStatus = "resolved"
	FraudStatusFalsePositive FraudStatus = "false_positive"
)

type FraudAlert struct {
	ID              uuid.UUID       `json:"id" db:"id"`
	UserID          uuid.UUID       `json:"user_id" db:"user_id"`
	AccountID       uuid.UUID       `json:"account_id" db:"account_id"`
	TransactionID   *uuid.UUID      `json:"transaction_id,omitempty" db:"transaction_id"`
	RuleName        string          `json:"rule_name" db:"rule_name"`
	Severity        FraudSeverity   `json:"severity" db:"severity"`
	Status          FraudStatus     `json:"status" db:"status"`
	RiskScore       int             `json:"risk_score" db:"risk_score"`
	Description     string          `json:"description" db:"description"`
	Details         json.RawMessage `json:"details,omitempty" db:"details"`
	ResolvedBy      *uuid.UUID      `json:"resolved_by,omitempty" db:"resolved_by"`
	ResolvedAt      *time.Time      `json:"resolved_at,omitempty" db:"resolved_at"`
	ResolutionNotes *string         `json:"resolution_notes,omitempty" db:"resolution_notes"`
	CreatedAt       time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at" db:"updated_at"`
}

type CreateFraudAlertRequest struct {
	UserID        uuid.UUID     `json:"user_id"`
	AccountID     uuid.UUID     `json:"account_id"`
	TransactionID *uuid.UUID    `json:"transaction_id,omitempty"`
	RuleName      string        `json:"rule_name"`
	Severity      FraudSeverity `json:"severity"`
	RiskScore     int           `json:"risk_score"`
	Description   string        `json:"description"`
	Details       interface{}   `json:"details,omitempty"`
}

type UpdateFraudAlertRequest struct {
	Status          *FraudStatus `json:"status,omitempty"`
	ResolvedBy      *uuid.UUID   `json:"resolved_by,omitempty"`
	ResolutionNotes *string      `json:"resolution_notes,omitempty"`
}

type FraudAlertFilter struct {
	UserID        *uuid.UUID     `json:"user_id,omitempty" query:"user_id"`
	AccountID     *uuid.UUID     `json:"account_id,omitempty" query:"account_id"`
	TransactionID *uuid.UUID     `json:"transaction_id,omitempty" query:"transaction_id"`
	RuleName      *string        `json:"rule_name,omitempty" query:"rule_name"`
	Severity      *FraudSeverity `json:"severity,omitempty" query:"severity"`
	Status        *FraudStatus   `json:"status,omitempty" query:"status"`
	MinRiskScore  *int           `json:"min_risk_score,omitempty" query:"min_risk_score"`
	MaxRiskScore  *int           `json:"max_risk_score,omitempty" query:"max_risk_score"`
	StartDate     *time.Time     `json:"start_date,omitempty" query:"start_date"`
	EndDate       *time.Time     `json:"end_date,omitempty" query:"end_date"`
	PaginationRequest
}

type FraudAlertResponse struct {
	ID              uuid.UUID            `json:"id"`
	UserID          uuid.UUID            `json:"user_id"`
	AccountID       uuid.UUID            `json:"account_id"`
	TransactionID   *uuid.UUID           `json:"transaction_id,omitempty"`
	RuleName        string               `json:"rule_name"`
	Severity        FraudSeverity        `json:"severity"`
	Status          FraudStatus          `json:"status"`
	RiskScore       int                  `json:"risk_score"`
	Description     string               `json:"description"`
	Details         json.RawMessage      `json:"details,omitempty"`
	ResolvedBy      *uuid.UUID           `json:"resolved_by,omitempty"`
	ResolvedAt      *time.Time           `json:"resolved_at,omitempty"`
	ResolutionNotes *string              `json:"resolution_notes,omitempty"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
	User            *UserProfile         `json:"user,omitempty"`
	Account         *AccountSummary      `json:"account,omitempty"`
	Transaction     *TransactionResponse `json:"transaction,omitempty"`
}
