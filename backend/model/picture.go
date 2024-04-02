package model

import "io"

type (
	Picture struct {
		ID          int64  `json:"id"`
		MemoryID    int64  `json:"memory_id"`
		PicturePath string `json:"Picture_path"`
	}

	CreatePictureRequest struct {
		MemoryID int64     `json:"memory_id"`
		Picture  io.Reader `json:"Picture"`
	}

	CreatePictureResponse struct {
		Picture Picture `json:"Picture"`
	}

	DeletePictureRequest struct {
		MemoryID  int64  `json:"memory_id"`
		PictureID string `json:"Picture_id"`
	}

	DeletePictureResponse struct{}
)
