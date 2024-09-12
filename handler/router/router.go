package router

import (
	"database/sql"
	"net/http"

	"github.com/sungyo4869/portfolio/handler"
	"github.com/sungyo4869/portfolio/handler/middleware"
	"github.com/sungyo4869/portfolio/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP)
	mux.HandleFunc("/user", handler.NewUserHandler(service.NewUserService(todoDB)).ServeHTTP)
	mux.HandleFunc("/login", handler.NewLoginHandler(service.NewUserService(todoDB)).ServeHTTP)
	mux.HandleFunc("/card", middleware.Auth(handler.NewCardHandler(service.NewCardService(todoDB))).ServeHTTP)
	
	return mux
}
