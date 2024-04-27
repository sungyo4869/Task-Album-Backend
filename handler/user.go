package handler

import "net/http"

type UserHandler struct{}

func NewUserinHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodGet:
	case http.MethodDelete:
	}
}
