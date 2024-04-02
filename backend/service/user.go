package service

import (
	"context"
	"database/sql"

	"github.com/sungyo4869/portfolio/model"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) CreateUser(ctx context.Context) (*model.User, error) {
	var user model.User

	return &user, nil
}
func (s *UserService) ReadUser(ctx context.Context) (*model.User, error) {
	var user model.User

	return &user, nil
}
func (s *UserService) UpdateUser(ctx context.Context) (*model.User, error) {
	var user model.User

	return &user, nil
}
func (s *UserService) DeleteUser(ctx context.Context) error {
	return nil
}

