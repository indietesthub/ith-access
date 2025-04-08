package repository

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/indietesthub/ith-access/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_GetByEmail(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				db: tt.fields.db,
			}
			got, err := r.GetByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetByID(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				db: tt.fields.db,
			}
			got, err := r.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	repo := NewMockUserRepository()

	// Test creating a new user
	newUser := &domain.User{
		Email:    "new@test.com",
		Password: "newPassword123",
		Name:     "New User",
	}

	err := repo.Create(newUser)
	assert.NoError(t, err)
	assert.NotZero(t, newUser.ID)
	assert.NotEmpty(t, newUser.Password)
	assert.NotEqual(t, "newPassword123", newUser.Password) // Password should be hashed
	assert.NotZero(t, newUser.CreatedAt)
	assert.NotZero(t, newUser.UpdatedAt)

	// Test creating user with existing email
	err = repo.Create(newUser)
	assert.Error(t, err)
	assert.Equal(t, ErrUserExists, err)
}

func TestGetUserByEmail(t *testing.T) {
	repo := NewMockUserRepository()

	// Test getting existing user
	user, err := repo.GetByEmail("test@test.com")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "test@test.com", user.Email)
	assert.Equal(t, "Test User", user.Name)

	// Test getting non-existing user
	user, err = repo.GetByEmail("nonexistent@test.com")
	assert.Error(t, err)
	assert.Equal(t, ErrUserNotFound, err)
	assert.Nil(t, user)
}
