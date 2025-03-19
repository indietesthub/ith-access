package database

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func RunMigrations() error {
	// First, verify we're connected to the right database
	var dbName string
	err := DB.QueryRow("SELECT current_database()").Scan(&dbName)
	if err != nil {
		return fmt.Errorf("error verifying database connection: %v", err)
	}

	files, err := os.ReadDir("internal/database/migrations")
	if err != nil {
		return fmt.Errorf("error reading migrations directory: %v", err)
	}

	// Sort files to ensure they run in order
	var migrationFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			migrationFiles = append(migrationFiles, file.Name())
		}
	}
	sort.Strings(migrationFiles)

	// Execute each migration in a transaction
	for _, file := range migrationFiles {
		content, err := os.ReadFile(filepath.Join("internal/database/migrations", file))
		if err != nil {
			return fmt.Errorf("error reading migration file %s: %v", file, err)
		}

		// Start a transaction
		tx, err := DB.Begin()
		if err != nil {
			return fmt.Errorf("error starting transaction for migration %s: %v", file, err)
		}

		// Execute the migration
		_, err = tx.Exec(string(content))
		if err != nil {
			// Rollback the transaction on error
			tx.Rollback()
			return fmt.Errorf("error executing migration %s: %v", file, err)
		}

		// Commit the transaction
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing migration %s: %v", file, err)
		}
	}

	return nil
}
