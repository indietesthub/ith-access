package auth

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSecretManager struct {
	mock.Mock
}

func (m *MockSecretManager) GetSecretValue(ctx context.Context, input *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*secretsmanager.GetSecretValueOutput), args.Error(1)
}

func MockGetJWTSecretKey(mockedSecret string) func() []byte {
	return func() []byte {
		return []byte(mockedSecret)
	}
}

func TestGenerateToken(t *testing.T) {
	// Save original JWTSecretKey
	originalJWTSecretKey := JWTSecretKey

	// Set test secret key
	JWTSecretKey = func() []byte {
		return []byte("test-secret-key")
	}

	// Restore original JWTSecretKey after test
	defer func() {
		JWTSecretKey = originalJWTSecretKey
	}()

	token, err := GenerateToken("user123")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken(t *testing.T) {
	// Save original JWTSecretKey
	originalJWTSecretKey := JWTSecretKey

	// Set test secret key
	JWTSecretKey = func() []byte {
		return []byte("test-secret-key")
	}

	// Restore original JWTSecretKey after test
	defer func() {
		JWTSecretKey = originalJWTSecretKey
	}()

	tokenString, err := GenerateToken("user123")
	assert.NoError(t, err)

	claims, err := ValidateToken(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, "user123", claims.UserID)
}

func TestValidateTokenInvalid(t *testing.T) {
	// Save original JWTSecretKey
	originalJWTSecretKey := JWTSecretKey

	// Set test secret key
	JWTSecretKey = func() []byte {
		return []byte("test-secret-key")
	}

	// Restore original JWTSecretKey after test
	defer func() {
		JWTSecretKey = originalJWTSecretKey
	}()

	_, err := ValidateToken("invalid-token")
	assert.Error(t, err)
}
