package repository

import (
	"time"

	"github.com/indietesthub/ith-access/internal/auth"
	"github.com/indietesthub/ith-access/internal/model"
)

type MockUserRepository struct {
	users map[string]*model.User
}

func NewMockUserRepository() *MockUserRepository {
	now := time.Now()

	// Create a hashed password for the test user
	hashedPassword, _ := auth.HashPassword("password")

	return &MockUserRepository{
		users: map[string]*model.User{
			"test@test.com": {
				ID:        1,
				Email:     "test@test.com",
				Password:  hashedPassword,
				Name:      "Test User",
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}
}

func (m *MockUserRepository) GetByEmail(email string) (*model.User, error) {
	if user, exists := m.users[email]; exists {
		return user, nil
	}
	return nil, ErrUserNotFound
}

func (m *MockUserRepository) GetByID(id int64) (*model.User, error) {
	for _, user := range m.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

func (m *MockUserRepository) Create(user *model.User) error {
	// Check if user already exists
	if _, exists := m.users[user.Email]; exists {
		return ErrUserExists
	}

	// Hash the password
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Set the hashed password
	user.Password = hashedPassword

	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Set ID (simulate auto-increment)
	user.ID = int64(len(m.users) + 1)

	// Store the user
	m.users[user.Email] = user
	return nil
}
