package repository

import (
	"fmt"
	"practice/pkg/models"
	"testing"
)

const connStr = "postgres://postgres:1@localhost:5432/postgres"

func TestBooksCRUD(t *testing.T) {
	db, err := New(connStr)
	if err != nil {
		t.Fatal(err)
	}

	book := models.Book{
		AuthorID: 3,
		GenreID:  4,
		Name:     "East of Eden",
	}

	id, err := db.NewBook(book)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("New Book was created at id: %d\n", id)

	_, err = db.GetBookByID(id)

	if err != nil {
		t.Fatal(err)
	}
}
