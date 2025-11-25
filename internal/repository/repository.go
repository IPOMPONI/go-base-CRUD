package repository

import (
	"context"

	"booklib/internal/model"
)

type BookRepo interface {
	InsertBook(ctx context.Context, book model.Book) error
	GetBookById(ctx context.Context, id int) (*model.Book, error)
	GetAllBooks(ctx context.Context) ([]model.Book, error)
	UpdateBookById(ctx context.Context, book model.Book) error
	DeleteBookById(ctx context.Context, id int) error
	DeleteAllBooks(ctx context.Context) error
}
