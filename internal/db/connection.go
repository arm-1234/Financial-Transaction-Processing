package db

import (
	"database/sql"
	"fmt"
	"time"

	"financial-transaction-system/internal/config"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Database struct {
	DB *sql.DB
}

func NewConnection(cfg *config.Config) (*Database, error) {
	db, err := sql.Open("postgres", cfg.GetDatabaseURL())
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logrus.Info("Successfully connected to database")
	return &Database{DB: db}, nil
}

func (d *Database) Close() error {
	if d.DB != nil {
		return d.DB.Close()
	}
	return nil
}

func (d *Database) Health() error {
	return d.DB.Ping()
}

// Transaction wrapper for database operations
func (d *Database) WithTransaction(fn func(*sql.Tx) error) error {
	tx, err := d.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p) // re-throw after rollback
		}
	}()

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			logrus.WithError(rbErr).Error("Failed to rollback transaction")
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
