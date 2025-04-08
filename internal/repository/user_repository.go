package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/indietesthub/ith-access/internal/auth"
	"github.com/indietesthub/ith-access/internal/database"
	"github.com/indietesthub/ith-access/internal/domain"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrDatabase     = errors.New("database error")
	ErrUserExists   = errors.New("user already exists")
)

type UserRepositoryInterface interface {
	GetByEmail(email string) (*domain.User, error)
	GetByID(id int64) (*domain.User, error)
	Create(user *domain.User) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.DB,
	}
}

const (
	selectUserFields = `
		SELECT id, email, password, name, created_at, updated_at 
		FROM users 
		WHERE %s = $1
	`
	insertUserFields = `
		INSERT INTO users (email, password, name, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id
	`
)

func (r *UserRepository) Create(user *domain.User) error {
	// Hash the password before storing
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDatabase, err)
	}

	// Check if user already exists
	existingUser, err := r.GetByEmail(user.Email)
	if err == nil && existingUser != nil {
		return ErrUserExists
	}

	// Insert the new user
	err = r.db.QueryRow(
		insertUserFields,
		user.Email,
		hashedPassword,
		user.Name,
	).Scan(&user.ID)

	if err != nil {
		return fmt.Errorf("%w: %v", ErrDatabase, err)
	}

	return nil
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	query := fmt.Sprintf(selectUserFields, "email")
	return r.getUser(query, email)
}

func (r *UserRepository) GetByID(id int64) (*domain.User, error) {
	query := fmt.Sprintf(selectUserFields, "id")
	return r.getUser(query, id)
}

func (r *UserRepository) getUser(query string, args ...interface{}) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow(query, args...).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabase, err)
	}

	return user, nil
}
