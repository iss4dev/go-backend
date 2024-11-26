package repository

import (
	"context"
	"practice/pkg/models"
)

func (repo *PGRepo) GetBooks() ([]models.Book, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	rows, err := repo.pool.Query(context.Background(),
		`SELECT id, name, author_id, genre_id FROM books;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.Book
	for rows.Next() {
		var item models.Book
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.AuthorID,
			&item.GenreID,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}

	return data, nil
}

func (repo *PGRepo) NewBook(item models.Book) (id int, err error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	err = repo.pool.QueryRow(context.Background(), `
		INSERT INTO books (name, author_id, genre_id)
		VALUES ($1, $2, $3) RETURNING id;`,
		item.Name,
		item.AuthorID,
		item.GenreID).Scan(&id)

	return id, err
}

func (repo *PGRepo) GetBookByID(id int) (models.Book, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var book models.Book
	err := repo.pool.QueryRow(context.Background(),
		`SELECT id, name, author_id, genre_id FROM books
		WHERE id = $1;`, id).Scan(
		&book.ID,
		&book.Name,
		&book.AuthorID,
		&book.GenreID,
	)

	return book, err
}
