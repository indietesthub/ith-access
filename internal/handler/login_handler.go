package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indietesthub/ith-access/internal/auth"
	"github.com/indietesthub/ith-access/internal/repository"
	"github.com/indietesthub/ith-access/pkg/logger"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("Invalid request body", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user from database
	userRepo := repository.NewUserRepository()
	user, err := userRepo.GetByEmail(request.Email)
	if err != nil {
		logger.Error("User not found or database error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// TODO: Add proper password hashing and comparison
	if user.Password != request.Password {
		logger.Error("Invalid password")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.Email)
	if err != nil {
		logger.Error("Failed to generate token", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	logger.Info("User logged in successfully")
	c.JSON(http.StatusOK, gin.H{"token": token})
}
