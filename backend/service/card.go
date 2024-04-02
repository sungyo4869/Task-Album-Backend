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

func (s *CardService) CreateCard(ctx context.Context) (*model.Card, error) {
	var card model.Card

	return &card, nil
}
func (s *CardService) ReadCard(ctx context.Context) ([]*model.Card, error) {
	var cards []*model.Card

	return cards, nil
}
func (s *CardService) UpdateCard(ctx context.Context) (*model.Card, error) {
	var card model.Card

	return &card, nil
}
func (s *CardService) DeleteCard(ctx context.Context) error {
	return nil
}

