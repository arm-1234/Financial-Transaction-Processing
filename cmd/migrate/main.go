package main

import (
	"fmt"
	"log"
	"os"

	"financial-transaction-system/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/migrate/main.go <up|down|force|version>")
	}

	command := os.Args[1]

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create migration instance
	m, err := migrate.New(
		"file://migrations",
		cfg.GetDatabaseURL(),
	)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}
	defer m.Close()

	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to run up migrations: %v", err)
		}
		fmt.Println("Migrations applied successfully")

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to run down migrations: %v", err)
		}
		fmt.Println("Migrations rolled back successfully")

	case "force":
		if len(os.Args) < 3 {
			log.Fatal("Usage: go run cmd/migrate/main.go force <version>")
		}
		version := os.Args[2]
		var v int
		if _, err := fmt.Sscanf(version, "%d", &v); err != nil {
			log.Fatalf("Invalid version: %v", err)
		}
		if err := m.Force(v); err != nil {
			log.Fatalf("Failed to force version: %v", err)
		}
		fmt.Printf("Forced database version to %d\n", v)

	case "version":
		version, dirty, err := m.Version()
		if err != nil {
			log.Fatalf("Failed to get version: %v", err)
		}
		fmt.Printf("Version: %d, Dirty: %t\n", version, dirty)

	default:
		log.Fatalf("Unknown command: %s. Use up, down, force, or version", command)
	}
}
