package model

type (
	Tag struct {
		ID       int64  `json:"id"`
		MemoryID int64  `json:"memory_id"`
		Label    string `json:"label"`
	}

	CreateTagRequest struct {
		MemoryID int64  `json:"memory_id"`
		Label    string `json:"label"`
	}

	CreateTagResponse struct {
		Tag Tag `json:"Tag"`
	}

	UpdateTagRequest struct {
		MemoryID int64  `json:"memory_id"`
		Label    string `json:"label"`
	}

	UpdateTagResponse struct {
		Tag Tag `json:"Tag"`
	}

	DeleteTagRequest struct {
		TagID int64 `json:"Tag_id"`
	}

	DeleteTagResponse struct{}
)
