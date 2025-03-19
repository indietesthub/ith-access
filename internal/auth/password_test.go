package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordHashing(t *testing.T) {
	password := "mySecurePassword123"

	// Test hashing
	hash, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, password, hash)

	// Test verification with correct password
	isValid := CheckPassword(password, hash)
	assert.True(t, isValid)

	// Test verification with wrong password
	isValid = CheckPassword("wrongPassword", hash)
	assert.False(t, isValid)
}

func TestPasswordHashingDifferentSalts(t *testing.T) {
	password := "mySecurePassword123"

	// Generate two hashes for the same password
	hash1, err1 := HashPassword(password)
	hash2, err2 := HashPassword(password)

	assert.NoError(t, err1)
	assert.NoError(t, err2)

	// Hashes should be different due to random salt
	assert.NotEqual(t, hash1, hash2)

	// Both should verify correctly
	assert.True(t, CheckPassword(password, hash1))
	assert.True(t, CheckPassword(password, hash2))
}
