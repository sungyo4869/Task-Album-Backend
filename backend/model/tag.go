package model

type (
	TAG struct{
		ID			int64	`json:"id"`
		MemoryID	int64	`json:"memory_id"`
		Label		string	`json:"label"`
	}
)
