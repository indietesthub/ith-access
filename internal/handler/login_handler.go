package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indietesthub/ith-access/internal/domain"
	"github.com/indietesthub/ith-access/pkg/logger"
)

type LoginHandler struct {
	usecase domain.LoginService
}

func NewLoginHandler(usecase domain.LoginService) *LoginHandler {
	return &LoginHandler{
		usecase: usecase,
	}
}

func (h *LoginHandler) Handle(c *gin.Context) {
	var request domain.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("Invalid request body", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user from database
	token, err := h.usecase.Login(&request)
	if err != nil {
		if err.Error() == "invalid credentials" {
			logger.Error("Invalid credentials", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		logger.Error("Login failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	logger.Info("User logged in successfully")
	c.JSON(http.StatusOK, gin.H{"token": token})
}
