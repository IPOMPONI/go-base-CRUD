package bookstorage

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertBook(conn *pgx.Conn, book Book) error {
	query := `INSERT INTO Book  (title, author, year_published) VALUES ($1, $2, $3)`

	_, err := conn.Exec(context.Background(), query, book.Title, book.Author, book.YearPublished)

	return err
}

func GetBookById(conn *pgx.Conn, id int) (*Book, error) {
	query := `SELECT id, title, author, year_published, added_at FROM Book WHERE id = $1`

	var book Book

	err := conn.QueryRow(context.Background(), query, id).Scan(
		&book.Id,
		&book.Title,
		&book.Author,
		&book.YearPublished,
		&book.AddedAt,
	)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func GetAllBooksData(conn *pgx.Conn) ([]Book, error) {
	query := `SELECT id, title, author, year_published, added_at FROM Book ORDER BY id`

	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bookAllData []Book

	for rows.Next() {
		var book Book

		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Author,
			&book.YearPublished,
			&book.AddedAt,
		)

		if err != nil {
			return nil, err
		}

		bookAllData = append(bookAllData, book)
	}

	return bookAllData, rows.Err()
}
