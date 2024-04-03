package service

import (
	"context"
	"database/sql"

	"github.com/sungyo4869/portfolio/model"
)

type CardService struct {
	db *sql.DB
}

func NewCardService(db *sql.DB) *CardService {
	return &CardService{
		db: db,
	}
}

func (s *CardService) CreateCard(ctx context.Context, title, summary, time_limit, status, description string) (*model.Card, error) {
	const (
		insert  = `INSERT INTO cards(title, summary, time_limit, status, description) VALUES(?, ?, ?, ?, ?)`
		confirm = `SELECT title, summary, time_limit, status, description FROM cards WHERE id = ?`
	)

	var card model.Card

	_, err := s.db.ExecContext(ctx, insert, title, summary, time_limit, status, description)

	if err != nil {
		return nil, err
	}

	return &card, nil
}
func (s *CardService) ReadCard(ctx context.Context) ([]*model.Card, error) {
	const (
		Read = `INSERT INTO cards(title, summary, time_limit, status, description) VALUES(?, ?, ?, ?,?) ORDER BY id DESC`
	)

	var cards []*model.Card

	return cards, nil
}
func (s *CardService) UpdateCard(ctx context.Context) (*model.Card, error) {
	const (
		update  = `UPDATE cards SET title = ?, summary = ?, time_limit = ?, status = ?, description = ? WHERE id = ?`
		confirm = `SELECT title, summary, time_limit, status, description FROM cards WHERE id = ?`
	)

	var card model.Card

	return &card, nil
}
func (s *CardService) DeleteCard(ctx context.Context) error {
	const deleteFmt = `DELETE FROM cards WHERE id IN (?%s)`

	return nil
}
