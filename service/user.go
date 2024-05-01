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

func (s *UserService) CreateUser(ctx context.Context, username, pass, email string) (*model.User, error) {
	const (
		insert  = `INSERT INTO users(username, password, email) VALUES(?, ?, ?)`
		confirm = `SELECT id, username, password, email FROM users WHERE id = ?`
	)

	result, err := s.db.ExecContext(ctx, insert, username, pass, email)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var user model.User
	row := s.db.QueryRowContext(ctx, confirm, id)
	if err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}
func (s *UserService) ReadUser(ctx context.Context, username string) (*model.User, error) {
	const query = `SELECT id, username, password, email FROM users WHERE username = ? AND password = ?`

	var user model.User
	row := s.db.QueryRowContext(ctx, query, username)
	if err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}
func (s *UserService) UpdateUser(ctx context.Context, username, pass, email string, id int64) (*model.User, error) {
	const (
		update  = `UPDATE users SET password = ?, email = ?, username = ? WHERE id = ?`
		confirm = `SELECT username, password, email FROM users WHERE id = ?`
	)

	_, err := s.db.ExecContext(ctx, update, pass, email, username, id)
	if err != nil {
		return nil, err
	}

	var user model.User
	row := s.db.QueryRowContext(ctx, confirm, id)

	err = row.Scan(
		&user.UserName, 
		&user.Password, 
		&user.Email,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (s *UserService) DeleteUser(ctx context.Context, username, pass string) error {
	const query = `DELETE FROM users WHERE username = ?`

	_, err := s.db.ExecContext(ctx, query, username)
	if err != nil {
		return err
	}

	return nil
}

