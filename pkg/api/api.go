package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	addr   string
	router *mux.Router
}

func New(addr string, router *mux.Router) *api {
	return &api{
		addr:   addr,
		router: router,
	}
}

func (api *api) FillEndpoints() {
	api.router.HandleFunc("/api/hello", helloHandler).Methods(http.MethodGet).Queries("name", "{name}")
	api.router.HandleFunc("/api/goodbye", goodbyeHandler).Methods(http.MethodGet)
}

func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.addr, api.router)
}
