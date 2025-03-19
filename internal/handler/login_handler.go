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

type LoginHandler struct {
	userRepo repository.UserRepositoryInterface
}

func NewLoginHandler(userRepo repository.UserRepositoryInterface) *LoginHandler {
	return &LoginHandler{
		userRepo: userRepo,
	}
}

func (h *LoginHandler) Handle(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("Invalid request body", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user from database
	user, err := h.userRepo.GetByEmail(request.Email)
	if err != nil {
		if err == repository.ErrUserNotFound {
			logger.Error("User not found")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		logger.Error("Database error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if user == nil {
		logger.Error("User not found")
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
