package model

type (
	User struct {
		ID       int64  `json:"id"`
		UserName string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	CreateUserRequest struct {
		UserName string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	CreateUserResponse struct {
		User User `json:"user"`
	}

	ReadUserRequest struct {
		Email string `json:"username"`
		Password string `json:"password"`
	}

	ReadUserResponse struct {
		User User `json:"user"`
	}

	UpdateUserRequest struct {
		UserName string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	UpdateUserResponse struct {
		User User `json:"user"`
	}

	DeleteUserRequest struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	DeleteUserResponse struct{}
)
