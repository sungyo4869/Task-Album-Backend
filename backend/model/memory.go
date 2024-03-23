package model

import "time"

type (
	MEMORY struct{
		ID			int64		`json:"id"`
		Title		string		`json:"title"`
		Summary		string		`json:"summary"`
		TimeLimit	time.Time	`json:"time_limit"`
		Status		string		`json:"status"`
		Description	string		`json:"description"`
	}
)
