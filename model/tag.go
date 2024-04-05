package model

type (
	Tag struct {
		ID     int64  `json:"id"`
		CardID int64  `json:"card_id"`
		Label  string `json:"label"`
	}

	CreateTagRequest struct {
		CardID int64  `json:"card_id"`
		Label  string `json:"label"`
	}

	CreateTagResponse struct {
		Tag Tag `json:"tag"`
	}

	UpdateTagRequest struct {
		CardID int64  `json:"card_id"`
		Label  string `json:"label"`
	}

	UpdateTagResponse struct {
		Tag Tag `json:"tag"`
	}

	DeleteTagRequest struct {
		TagID int64 `json:"tag_id"`
	}

	DeleteTagResponse struct{}
)
