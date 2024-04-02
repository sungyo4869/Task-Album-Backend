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

	CreateTaskRequest struct {
		Title       string    `json:"title"`
		Summary     string    `json:"summary"`
		TimeLimit   time.Time `json:"time_limit"`
		Status      string    `json:"status"`
		Description string    `json:"description"`
	}

	CreateTaskResponse struct {
		Task Card `json:"task"`
	}

	ReadTasksRequest struct {
		PrevID int64 `json:"prev_id"`
	}

	ReadTasksResponse struct {
		Memories []Card `json:"memories"`
	}

	UpdateTaskRequest struct {
		CardID    int64     `json:"Card_id"`
		Title       string    `json:"title"`
		Summary     string    `json:"summary"`
		TimeLimit   time.Time `json:"time_limit"`
		Description string    `json:"description"`
	}

	UpdateTaskResponse struct {
		Task Card `json:"task"`
	}

	UpdateStatusRequest struct {
		CardID int64  `json:"Card_id"`
		Status   string `json:"status"`
	}

	UpdateStatusResponse struct {
		Card Card `json:"Card"`
	}

	DeleteCardRequest struct {
		CardID Card `json:"Card"`
	}

	DeleteCardResponse struct{}
)
