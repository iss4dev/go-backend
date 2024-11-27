package models

type Book struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	AuthorID int    `json:"author_id"`
	GenreID  int    `json:"genre_id"`
}

type Genre struct {
	ID    int
	Genre string
}

type Author struct {
	ID   int
	Name string
}
