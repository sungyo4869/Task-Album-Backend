package model

import "io"

type (
	PICTURE struct {
		ID          int64  `json:"id"`
		MemoryID    int64  `json:"memory_id"`
		PicturePath string `json:"picture_path"`
	}

	CreatePictureRequest struct {
		MemoryID int64     `json:"memory_id"`
		Picture  io.Reader `json:"picture"`
	}

	CreatePictureResponse struct {
		Picture PICTURE `json:"picture"`
	}

	DeletePictureRequest struct {
		MemoryID  int64  `json:"memory_id"`
		PictureID string `json:"picture_id"`
	}

	DeletePictureResponse struct{}
)
