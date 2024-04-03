package handler

import "net/http"

type UserHandler struct{}

func NewUserinHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
