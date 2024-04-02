package model

import "io"

type (
	Picture struct {
		ID          int64  `json:"id"`
		CardID    int64  `json:"card_id"`
		PicturePath string `json:"picture_path"`
	}

	CreatePictureRequest struct {
		CardID int64     `json:"card_id"`
		Picture  io.Reader `json:"picture"`
	}

	CreatePictureResponse struct {
		Picture Picture `json:"picture"`
	}

	DeletePictureRequest struct {
		CardID  int64  `json:"card_id"`
		PictureID string `json:"picture_id"`
	}

	DeletePictureResponse struct{}
)
