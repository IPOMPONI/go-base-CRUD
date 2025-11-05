package repository

import (
	"context"

	"booklib/internal/model"

	"github.com/jackc/pgx/v5"
)

type PgRepo struct {
	conn *pgx.Conn
}

func NewBookRepo(conn *pgx.Conn) *PgRepo {
	return &PgRepo{conn: conn}
}

func (pgr *PgRepo) InsertBook(ctx context.Context, book model.Book) error {
	query := `INSERT INTO Books (title, author, year_published) VALUES ($1, $2, $3)`

	_, err := pgr.conn.Exec(ctx, query, book.Title, book.Author, book.YearPublished)

	return err
}

func (pgr *PgRepo) GetBookById(ctx context.Context, id int) (*model.Book, error) {
	query := `SELECT id, title, author, year_published, added_at FROM Books WHERE id = $1`

	var book model.Book

	err := pgr.conn.QueryRow(ctx, query, id).Scan(
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

func (pgr *PgRepo) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	query := `SELECT id, title, author, year_published, added_at FROM Books ORDER BY id`

	rows, err := pgr.conn.Query(ctx, query)

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

func (pgr *PgRepo) UpdateBookById(ctx context.Context, book model.Book) error {
	query := `UPDATE Books SET title = $1, author = $2, year_published = $3 WHERE id = $4`

	_, err := pgr.conn.Exec(ctx, query, book.Title, book.Author, book.YearPublished, book.Id)

	return err
}

func (pgr *PgRepo) DeleteBookById(ctx context.Context, id int) error {
	query := `DELETE FROM Books WHERE id = $1`

	_, err := pgr.conn.Exec(ctx, query, id)

	return err
}

func (pgr *PgRepo) DeleteAllBooks(ctx context.Context) error {
	query := `DELETE FROM Books`

	_, err := pgr.conn.Exec(ctx, query)

	return err
}
