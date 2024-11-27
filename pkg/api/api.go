package api

import (
	"net/http"
	"practice/pkg/repository"

	"github.com/gorilla/mux"
)

type api struct {
	addr   string
	router *mux.Router
	db     *repository.PGRepo
}

func New(addr string, router *mux.Router, db *repository.PGRepo) *api {
	return &api{
		addr:   addr,
		router: router,
		db:     db,
	}
}

func (api *api) FillEndpoints() {
	api.router.HandleFunc("/api/books", api.booksHandler).Queries("id", "{id}")
	api.router.HandleFunc("/api/books", api.booksHandler)
}

func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.addr, api.router)
}
