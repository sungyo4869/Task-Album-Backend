package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sungyo4869/portfolio/model"
	"github.com/sungyo4869/portfolio/service"
)

type UserHandler struct{
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var req model.ReadUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Error Decoding Request", http.StatusBadRequest)
			return
		}

		res, err := u.Read(r.Context(), &req)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			http.Error(w, "Error Encoding Response", http.StatusInternalServerError)
		}
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	}
}

func (h *UserHandler) Create(ctx context.Context, req *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	var res model.CreateUserResponse

	result, err := h.svc.CreateUser(ctx, req.UserName, req.Password, req.Email)
	if err != nil {
		return nil, err
	}

	res.User = *result

	return &res, nil
}

func (h *UserHandler) Read(ctx context.Context, req *model.ReadUserRequest) (*model.ReadUserResponse, error) {
	var res model.ReadUserResponse

	result, err := h.svc.ReadUser(ctx, req.UserName, req.Password)
	if err != nil {
		return nil, err
	}

	res.User = *result

	return &res, nil
}

func (h *UserHandler) Update(ctx context.Context, req *model.UpdateUserRequest, id int64) (*model.UpdateUserResponse, error) {
	var res model.UpdateUserResponse

	result, err := h.svc.UpdateUser(ctx, req.UserName, req.Password, req.Email, id)
	if err != nil {
		return nil, err
	}

	res.User = *result

	return &res, nil
}

func (h *UserHandler) Delete(ctx context.Context, req *model.DeleteUserRequest) (*model.DeleteUserResponse, error) {
	var res model.DeleteUserResponse

	err := h.svc.DeleteUser(ctx, req.UserName, req.Password)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
