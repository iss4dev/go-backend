package main

import (
	"log"
	"practice/pkg/api"
	"practice/pkg/repository"

	"github.com/gorilla/mux"
)

const connStr = "postgres://postgres:psql@localhost:5432/postgres"

func main() {
	db, err := repository.New(connStr)
	if err != nil {
		log.Fatal(err)
	}

	api := api.New("localhost:8090", mux.NewRouter(), db)
	api.FillEndpoints()
	api.ListenAndServe()

	// pg, err := repository.New(connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// books, err := pg.GetBooks()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, book := range books {
	// 	fmt.Printf("%d) Book name: %s Author ID: %d enre ID: %d\n", book.ID, book.Name, book.AuthorID, book.GenreID)
	// }
}
