package postgresql

import (
	"time"
	"context"
	"testing"

	"booklib/internal/domain"
)

func TestInsertBook(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, _ := NewConnectDb(ctx)
	repo := NewBookRepo(db)

	_, err := db.Exec(ctx, "TRUNCATE TABLE books RESTART IDENTITY")
	if err != nil {
		t.Fatalf("Error truncate table! %v", err)
	}

	defer db.Close(ctx)

	book := domain.Book{
		Title:         "Книга 1",
		Author:        "Автор 1",
		YearPublished: 2025,
	}

	err = repo.InsertBook(ctx, book)
	if err != nil {
		t.Fatalf("Error insert book! %v", err)
	}
}

func TestInsertBookExistingTitle(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, _ := NewConnectDb(ctx)
	repo := NewBookRepo(db)

	defer db.Close(ctx)

	book := domain.Book{
		Title:         "Книга 1",
		Author:        "Автор 1",
		YearPublished: 2025,
	}

	err := repo.InsertBook(ctx, book)
	if err == nil {
		t.Error("Expected error for existing title")
	}
}

func TestInsertBookZeroYear(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, _ := NewConnectDb(ctx)
	repo := NewBookRepo(db)

	defer db.Close(ctx)

	book := domain.Book{
		Title:         "Книга 1",
		Author:        "Автор 1",
		YearPublished: 0,
	}

	err := repo.InsertBook(ctx, book)
	if err == nil {
		t.Error("Expected error for zero year.")
	}
}

func TestInsertBookFutureYear(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, _ := NewConnectDb(ctx)
	repo := NewBookRepo(db)

	defer db.Close(ctx)

	book := domain.Book{
		Title:         "Книга 1",
		Author:        "Автор 1",
		YearPublished: 2077,
	}

	err := repo.InsertBook(ctx, book)
	if err == nil {
		t.Error("Expected error for future year.")
	}
}

func TestGetBookById(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, _ := NewConnectDb(ctx)
	repo := NewBookRepo(db)

	defer db.Close(ctx)

	book, err := repo.GetBookById(ctx, 1)
	if book == nil {
		t.Fatalf("Geted book is nil!")
	}

	if err != nil {
		t.Fatalf("Error get book by id! %v", err)
	}
}

func TestGetBookByIncorrectId(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, _ := NewConnectDb(ctx)
	repo := NewBookRepo(db)

	defer db.Close(ctx)

	book, err := repo.GetBookById(ctx, 2)
	if book != nil {
		t.Error("Expected to be nil for incorrect id.")
	}

	if err == nil {
		t.Error("Expected error for incorrect id.")
	}
}

func TestGetAllBooks(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, _ := NewConnectDb(ctx)
	repo := NewBookRepo(db)

	defer db.Close(ctx)

	books, err := repo.GetAllBooks(ctx)
	if len(books) == 0 {
		t.Fatal("Getted data is empty")
	}

	if err != nil {
		t.Fatalf("Error getted all books! %v", err)
	}
}

func TestUpdateBookById(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, _ := NewConnectDb(ctx)
	repo := NewBookRepo(db)

	defer db.Close(ctx)

	book := domain.Book{
		Id:            1,
		Title:         "Новое название",
		Author:        "Новый автор",
		YearPublished: 2024,
	}

	err := repo.UpdateBookById(ctx, book)
	if err != nil {
		t.Fatalf("Error update book! %v", err)
	}
}

func TestDeleteBookById(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, _ := NewConnectDb(ctx)
	repo := NewBookRepo(db)

	defer db.Close(ctx)

	err := repo.DeleteBookById(ctx, 1)
	if err != nil {
		t.Fatalf("Error delete book! %v", err)
	}
}

func TestDeleteAllBooks(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, _ := NewConnectDb(ctx)
	repo := NewBookRepo(db)

	defer db.Close(ctx)

	err := repo.DeleteAllBooks(ctx)
	if err != nil {
		t.Fatalf("Error delete all books! %v", err)
	}
}
