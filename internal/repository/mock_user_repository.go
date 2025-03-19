package repository

import (
	"github.com/indietesthub/ith-access/internal/model"
)

type MockUserRepository struct {
	users map[string]*model.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: map[string]*model.User{
			"test@test.com": {
				ID:       1,
				Email:    "test@test.com",
				Password: "password",
				Name:     "Test User",
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
