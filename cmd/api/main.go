package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/indietesthub/ith-access/internal/database"
	"github.com/indietesthub/ith-access/internal/handler"
	"github.com/indietesthub/ith-access/internal/repository"
	"github.com/indietesthub/ith-access/pkg/logger"
)

func main() {
	// Initialize logger
	logger.InitLogger()

	// Initialize database
	logger.Info("Initializing database")
	err := database.InitDB()
	if err != nil {
		logger.Fatal("Failed to initialize database:", err)
	}

	// Run migrations
	err = database.RunMigrations()
	if err != nil {
		logger.Fatal("Failed to run migrations:", err)
	}

	// Set up router
	router := gin.Default()

	// Add routes
	userRepo := repository.NewUserRepository()
	loginHandler := handler.NewLoginHandler(userRepo)
	router.POST("/login", loginHandler.Handle)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Starting server on port", port)
	if err := router.Run(":" + port); err != nil {
		logger.Fatal("Failed to start server:", err)
	}
}
