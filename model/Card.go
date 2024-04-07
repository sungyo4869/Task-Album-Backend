package model

import "time"

type (
	Card struct {
		ID          int64     `json:"id"`
		Title       string    `json:"title"`
		Summary     string    `json:"summary"`
		TimeLimit   time.Time `json:"time_limit"`
		Status      string    `json:"status"`
		Description string    `json:"description"`
	}

	CreateCardRequest struct {
		Title       string    `json:"title"`
		Summary     string    `json:"summary"`
		TimeLimit   time.Time `json:"time_limit"`
		Description string    `json:"description"`
	}

	CreateCardResponse struct {
		Card Card `json:"card"`
	}

	ReadCardsRequest struct {}

	ReadCardsResponse struct {
		Memories []*Card `json:"memories"`
	}

	UpdateCardRequest struct {
		CardID    int64     `json:"Card_id"`
		Title       string    `json:"title"`
		Summary     string    `json:"summary"`
		TimeLimit   time.Time `json:"time_limit"`
		Description string    `json:"description"`
	}

	UpdateCardResponse struct {
		Card Card `json:"card"`
	}

	UpdateStatusRequest struct {
		CardID int64  `json:"card_id"`
		Status   string `json:"status"`
	}

	UpdateStatusResponse struct {
		Card Card `json:"card"`
	}

	DeleteCardRequest struct {
		CardID int64 `json:"card_id"`
	}

	DeleteCardResponse struct{}
)
