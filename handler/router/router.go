package router

import (
	"database/sql"
	"net/http"

	"github.com/sungyo4869/portfolio/handler"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP)

	return mux
}
