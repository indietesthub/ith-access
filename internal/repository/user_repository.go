package repository

import (
	"database/sql"
	"fmt"

	"github.com/indietesthub/ith-access/internal/database"
	"github.com/indietesthub/ith-access/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.DB,
	}
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}

	query := `
		SELECT id, email, password, name, created_at, updated_at 
		FROM users 
		WHERE email = ?
	`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error querying user: %v", err)
	}

	return user, nil
}

func (r *UserRepository) GetByID(id int64) (*model.User, error) {
	user := &model.User{}

	query := `
		SELECT id, email, password, name, created_at, updated_at 
		FROM users 
		WHERE id = ?
	`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error querying user: %v", err)
	}

	return user, nil
}
