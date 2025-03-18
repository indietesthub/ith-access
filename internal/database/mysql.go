package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/indietesthub/ith-access/pkg/logger"
)

var DB *sql.DB

func InitDB() error {
	logger.Info("Getting database credentials")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	logger.Info("Connecting to database")
	var err error
	DB, err = sql.Open("mysql", dsn)
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
