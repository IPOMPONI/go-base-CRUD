package domain

import (
	"context"
	"time"
)

type Book struct {
	Id            int       `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	YearPublished int       `json:"year_published"`
	AddedAt       time.Time `json:"added_at"`
}

type BookRepo interface {
	InsertBook(ctx context.Context, book Book) error
	GetBookById(ctx context.Context, id int) (*Book, error)
	GetAllBooks(ctx context.Context) ([]Book, error)
	UpdateBookById(ctx context.Context, book Book) error
	DeleteBookById(ctx context.Context, id int) error
	DeleteAllBooks(ctx context.Context) error
}
