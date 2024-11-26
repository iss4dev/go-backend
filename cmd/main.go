package main

import (
	"practice/pkg/api"

	"github.com/gorilla/mux"
)

// const connStr = "postgres://postgres:1@localhost:5432/postgres"

func main() {
	api := api.New("localhost:8090", mux.NewRouter())
	api.FillEndpoints()
	api.ListenAndServe()
}
