package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/indietesthub/ith-access/internal/config"
	"github.com/indietesthub/ith-access/internal/database"
	"github.com/indietesthub/ith-access/internal/handler"
	"github.com/indietesthub/ith-access/internal/middleware"
	"github.com/indietesthub/ith-access/internal/repository"
	"github.com/indietesthub/ith-access/pkg/logger"
)

func main() {
	// Initialize logger
	logger.InitLogger()

	// Load configuration
	logger.Info("Loading configuration")
	_, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load configuration:", err)
	}

	// Initialize database
	logger.Info("Initializing database")
	err = database.InitDB()
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

	// Add security middlewares
	router.Use(middleware.SecurityHeaders())
	router.Use(middleware.CORSMiddleware())

	// Add routes with API key protection
	userRepo := repository.NewUserRepository()
	loginHandler := handler.NewLoginHandler(userRepo)
	signupHandler := handler.NewSignupHandler(userRepo)

	// Public routes with API key protection
	public := router.Group("/api/v1")
	public.Use(middleware.RequireAPIKey())
	public.Use(middleware.RateLimit())
	{
		public.POST("/login", loginHandler.Handle)
		public.POST("/signup", signupHandler.Handle)
	}

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
