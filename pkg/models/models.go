package models

type Book struct {
	ID       int
	Name     string
	AuthorID int
	GenreID  int
}

type Genre struct {
	ID    int
	Genre string
}

type Author struct {
	ID   int
	Name string
}
