package repository

import (
	"database/sql"
	"errors"

	"github.com/indietesthub/ith-access/internal/database"
	"github.com/indietesthub/ith-access/internal/model"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrDatabase     = errors.New("database error")
)

type UserRepositoryInterface interface {
	GetByEmail(email string) (*model.User, error)
	GetByID(id int64) (*model.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.DB,
	}
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, email, password, name FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, ErrDatabase
	}
	return &user, nil
}

func (r *UserRepository) GetByID(id int64) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, email, password, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, ErrDatabase
	}
	return &user, nil
}
