package model

import "time"

type (
	MEMORY struct {
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
		Task MEMORY `json:"task"`
	}

	ReadTasksRequest struct {
		PrevID int64 `json:"prev_id"`
	}

	ReadTasksResponse struct {
		Memories []MEMORY `json:"memories"`
	}

	UpdateTaskRequest struct {
		MemoryID    int64     `json:"memory_id"`
		Title       string    `json:"title"`
		Summary     string    `json:"summary"`
		TimeLimit   time.Time `json:"time_limit"`
		Description string    `json:"description"`
	}

	UpdateTaskResponse struct {
		Task MEMORY `json:"task"`
	}

	UpdateStatusRequest struct {
		MemoryID int64  `json:"memory_id"`
		Status   string `json:"status"`
	}

	UpdateStatusResponse struct {
		Memory MEMORY `json:"memory"`
	}

	DeleteMemoryRequest struct {
		MemoryID MEMORY `json:"memory"`
	}

	DeleteMemoryResponse struct{}
)
