package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indietesthub/ith-access/internal/model"
	"github.com/indietesthub/ith-access/internal/repository"
	"github.com/indietesthub/ith-access/pkg/logger"
)

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type SignupHandler struct {
	userRepo repository.UserRepositoryInterface
}

func NewSignupHandler(userRepo repository.UserRepositoryInterface) *SignupHandler {
	return &SignupHandler{
		userRepo: userRepo,
	}
}

func (h *SignupHandler) Handle(c *gin.Context) {
	var request SignupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("Invalid request body", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	existingUser, err := h.userRepo.GetByEmail(request.Email)
	if err == nil && existingUser != nil {
		logger.Error("User already exists")
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Create new user
	newUser := &model.User{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
	}

	err = h.userRepo.Create(newUser)
	if err != nil {
		logger.Error("Failed to create user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	logger.Info("User created successfully")
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
