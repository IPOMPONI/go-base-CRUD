package repository

import (
	"context"

	"booklib/internal/model"

	"github.com/jackc/pgx/v5"
)

func InsertBook(conn *pgx.Conn, book model.Book) error {
	query := `INSERT INTO Books (title, author, year_published) VALUES ($1, $2, $3)`

	_, err := conn.Exec(context.Background(), query, book.Title, book.Author, book.YearPublished)

	return err
}

func GetBookById(conn *pgx.Conn, id int) (*model.Book, error) {
	query := `SELECT id, title, author, year_published, added_at FROM Books WHERE id = $1`

	var book model.Book

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

func GetAllBooks(conn *pgx.Conn) ([]model.Book, error) {
	query := `SELECT id, title, author, year_published, added_at FROM Books ORDER BY id`

	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bookAllData []model.Book

	for rows.Next() {
		var book model.Book

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

func UpdateBookById(conn *pgx.Conn, book model.Book) error {
	query := `UPDATE Books SET title = $1, author = $2, year_published = $3 WHERE id = $4`

	_, err := conn.Exec(context.Background(), query, book.Title, book.Author, book.YearPublished, book.Id)

	return err
}

func DeleteBookById(conn *pgx.Conn, id int) error {
	query := `DELETE FROM Books WHERE id = $1`

	_, err := conn.Exec(context.Background(), query, id)

	return err
}

func DeleteAllBooks(conn *pgx.Conn) error {
	query := `DELETE FROM Books`

	_, err := conn.Exec(context.Background(), query)

	return err
}
