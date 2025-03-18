package database

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func RunMigrations() error {
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

	// Execute each migration
	for _, file := range migrationFiles {
		content, err := os.ReadFile(filepath.Join("internal/database/migrations", file))
		if err != nil {
			return fmt.Errorf("error reading migration file %s: %v", file, err)
		}

		_, err = DB.Exec(string(content))
		if err != nil {
			return fmt.Errorf("error executing migration %s: %v", file, err)
		}
	}

	return nil
}
