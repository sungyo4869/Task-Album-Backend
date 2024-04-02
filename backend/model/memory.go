package model

import "time"

type (
	Memory struct {
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
		Task Memory `json:"task"`
	}

	ReadTasksRequest struct {
		PrevID int64 `json:"prev_id"`
	}

	ReadTasksResponse struct {
		Memories []Memory `json:"memories"`
	}

	UpdateTaskRequest struct {
		MemoryID    int64     `json:"Memory_id"`
		Title       string    `json:"title"`
		Summary     string    `json:"summary"`
		TimeLimit   time.Time `json:"time_limit"`
		Description string    `json:"description"`
	}

	UpdateTaskResponse struct {
		Task Memory `json:"task"`
	}

	UpdateStatusRequest struct {
		MemoryID int64  `json:"Memory_id"`
		Status   string `json:"status"`
	}

	UpdateStatusResponse struct {
		Memory Memory `json:"Memory"`
	}

	DeleteMemoryRequest struct {
		MemoryID Memory `json:"Memory"`
	}

	DeleteMemoryResponse struct{}
)
