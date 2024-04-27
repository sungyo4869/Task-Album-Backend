package service

import (
	"context"
	"database/sql"
	"time"

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

func (s *CardService) CreateCard(ctx context.Context, title, summary string,  timeLimit time.Time) (*model.Card, error) {
	const (
		insert  = `INSERT INTO cards(title, summary, time_limit) VALUES(?, ?, ?, ?)`
		confirm = `SELECT title, summary, time_limit, status, FROM cards WHERE id = ?`
	)

	var card model.Card

	result, err := s.db.ExecContext(ctx, insert, title, summary, timeLimit)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, confirm, id)

	err = row.Scan(
		&card.Title,
		&card.Summary,
		&card.TimeLimit,
		&card.Status,
	)

	card.ID = id

	if err != nil {
		return nil, err
	}

	return &card, nil
}
func (s *CardService) ReadCard(ctx context.Context) ([]*model.Card, error) {
	const (
		Read =`SELECT title, summary, time_limit, status, description FROM cards WHERE id = ?`
	)
	var cards []*model.Card

	rows, err := s.db.QueryContext(ctx, Read)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var card model.Card
		err := rows.Scan(
			&card.ID,
			&card.Title,
			&card.Summary,
			&card.TimeLimit,
			&card.Status,
			&card.Description,
		)
		if err != nil {
			return nil, err
		}
		cards = append(cards, &card)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}
func (s *CardService) UpdateCard(ctx context.Context, title, summary, description string, timeLimit time.Time, id int64) (*model.Card, error) {
	const (
		update  = `UPDATE cards SET title = ?, summary = ?, time_limit = ?, description = ? WHERE id = ?`
		confirm = `SELECT title, summary, time_limit, status, description FROM cards WHERE id = ?`
	)

	var card model.Card

	_, err := s.db.ExecContext(ctx, update, title, summary, timeLimit, description, id)
	if err != nil {
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, confirm)

	err = row.Scan(
		&card.ID,
		&card.Title,
		&card.Summary,
		&card.TimeLimit,
		&card.Status,
		&card.Description,
	)

	if err != nil {
		return nil, err
	}

	return &card, nil
}
func (s *CardService) DeleteCard(ctx context.Context, id int64) error {
	const delete = `DELETE FROM cards WHERE id IN (?%s)`

	_, err := s.db.ExecContext(ctx, delete, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *CardService) UpdateStatus(ctx context.Context, id int, status string) (*model.Card, error) {
	const (
		update = `UPDATE cards SET status = ? WHERE id = ?`
		confirm = `SELECT title, summary, time_limit, status, description FROM cards WHERE id = ?`
	)

	var card model.Card

	_, err := s.db.ExecContext(ctx, update, status, id)
	if err != nil {
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, confirm)

	err = row.Scan(
		&card.ID,
		&card.Title,
		&card.Summary,
		&card.TimeLimit,
		&card.Status,
		&card.Description,
	)

	if err != nil {
		return nil, err
	}

	return &card, nil
}
