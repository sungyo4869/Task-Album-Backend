package model

type (
	TAG struct {
		ID       int64  `json:"id"`
		MemoryID int64  `json:"memory_id"`
		Label    string `json:"label"`
	}

	CreateTagRequest struct {
		MemoryID int64  `json:"memory_id"`
		Label    string `json:"label"`
	}

	CreateTagResponse struct {
		Tag TAG `json:"tag"`
	}

	UpdateTagRequest struct {
		MemoryID int64  `json:"memory_id"`
		Label    string `json:"label"`
	}

	UpdateTagResponse struct {
		Tag TAG `json:"tag"`
	}

	DeleteTagRequest struct {
		TagID int64 `json:"tag_id"`
	}

	DeleteTagResponse struct{}
)
