package model

type (
	PICTURE struct{
		ID				int64	`json:"id"`
		MemoryID		int64	`json:"memory_id"`
		PicturePath		string	`json:"picture_path"`
	}
)
