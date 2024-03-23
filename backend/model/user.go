package model

type (
	USER struct{
		ID			int64	`json:"id"`
		UserName	string	`json:"username"`
		Password	string	`json:"password"`
		Email		string	`json:"email"`
	}
)
