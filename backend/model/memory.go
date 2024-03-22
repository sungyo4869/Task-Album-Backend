package model

import "time"

type (
	MEMORY struct{
		ID			int64
		Title		string
		Summary		string
		TimeLimit	time.Time
		Status		string
		Description	string
	}
)
