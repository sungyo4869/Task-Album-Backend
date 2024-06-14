package model

type LoginRequest struct {
	Email string `json:"email"`
	Pass string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"jwt-token"`
}

type LogoutRequest struct {
	Email string `json:"email"`
	Pass string `json:"password"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}
