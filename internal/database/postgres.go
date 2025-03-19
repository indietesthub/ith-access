package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/indietesthub/ith-access/pkg/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	// Load .env file from the root of the project
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	logger.Info("Getting database credentials")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Validate that all required environment variables are set
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return fmt.Errorf("missing required database environment variables")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	logger.Info("Connecting to database")
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	logger.Info("Testing connection to database")
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	logger.Info("Database connection successful")
	return nil
}
