package model

type (
	User struct {
		ID       int64  `json:"id"`
		UserName string `json:"Username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	CreateUserRequest struct {
		UserName string `json:"Username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	CreateUserResponse struct {
		User User `json:"User"`
	}

	ReadUserRequest struct {
		UserName string `json:"Username"`
	}

	ReadUserResponse struct {
		User User `json:"User"`
	}

	UpdateUserRequest struct {
		UserName string `json:"Username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	UpdateUserResponse struct {
		User User `json:"User"`
	}

	DeleteUserRequest struct {
		UserName string `json:"Username"`
		Password string `json:"password"`
	}

	DeleteUserResponse struct{}
)
