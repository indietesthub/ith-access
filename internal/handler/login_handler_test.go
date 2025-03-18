package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/login", LoginHandler)
	return r
}

func TestLoginSuccess(t *testing.T) {
	router := setupRouter()

	loginRequest := LoginRequest{
		Email:    "test@test.com",
		Password: "password",
	}
	jsonData, _ := json.Marshal(loginRequest)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "token")
	assert.NotEmpty(t, response["token"])
}

func TestLoginInvalidCredentials(t *testing.T) {
	router := setupRouter()

	loginRequest := LoginRequest{
		Email:    "wrong@test.com",
		Password: "wrongpassword",
	}
	jsonData, _ := json.Marshal(loginRequest)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid credentials", response["error"])
}

func TestLoginInvalidRequest(t *testing.T) {
	router := setupRouter()

	// Missing required fields
	loginRequest := map[string]string{
		"email": "test@test.com",
	}
	jsonData, _ := json.Marshal(loginRequest)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
