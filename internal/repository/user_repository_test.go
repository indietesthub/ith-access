package repository

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/indietesthub/ith-access/internal/model"
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
		want    *model.User
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
		want    *model.User
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
