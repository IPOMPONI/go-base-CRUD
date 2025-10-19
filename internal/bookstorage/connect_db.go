package bookstorage

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewConnectDb() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), "user=postgres dbname=books_db sslmode=disable")
}
