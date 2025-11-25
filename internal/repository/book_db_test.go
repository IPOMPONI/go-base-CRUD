package repository

import (
	"context"
	"testing"

	"booklib/internal/domain"
)

func TestInsertBook(t *testing.T) {
	db, _ := NewConnectDb()
	repo := NewBookRepo(db)

	_, err := db.Exec(context.Background(), "TRUNCATE TABLE books RESTART IDENTITY")
	if err != nil {
		t.Fatalf("Error truncate table!")
	}

	defer db.Close(context.Background())

	book := domain.Book{
		Title:         "Книга 1",
		Author:        "Автор 1",
		YearPublished: 2025,
	}

	err = repo.InsertBook(context.Background(), book)
	if err != nil {
		t.Fatalf("Error insert book! %v", err)
	}
}

func TestInsertBookExistingTitle(t *testing.T) {
	db, _ := NewConnectDb()
	repo := NewBookRepo(db)

	defer db.Close(context.Background())

	book := domain.Book{
		Title:         "Книга 1",
		Author:        "Автор 1",
		YearPublished: 2025,
	}

	err := repo.InsertBook(context.Background(), book)
	if err == nil {
		t.Error("Expected error for existing title")
	}
}

func TestInsertBookZeroYear(t *testing.T) {
	db, _ := NewConnectDb()
	repo := NewBookRepo(db)

	defer db.Close(context.Background())

	book := domain.Book{
		Title:         "Книга 1",
		Author:        "Автор 1",
		YearPublished: 0,
	}

	err := repo.InsertBook(context.Background(), book)
	if err == nil {
		t.Error("Expected error for zero year.")
	}
}

func TestInsertBookFutureYear(t *testing.T) {
	db, _ := NewConnectDb()
	repo := NewBookRepo(db)

	defer db.Close(context.Background())

	book := domain.Book{
		Title:         "Книга 1",
		Author:        "Автор 1",
		YearPublished: 2077,
	}

	err := repo.InsertBook(context.Background(), book)
	if err == nil {
		t.Error("Expected error for future year.")
	}
}

func TestGetBookById(t *testing.T) {
	db, _ := NewConnectDb()
	repo := NewBookRepo(db)

	defer db.Close(context.Background())

	book, err := repo.GetBookById(context.Background(), 1)
	if book == nil {
		t.Fatalf("Geted book is nil!")
	}

	if err != nil {
		t.Fatalf("Error get book by id! %v", err)
	}
}

func TestGetBookByIncorrectId(t *testing.T) {
	db, _ := NewConnectDb()
	repo := NewBookRepo(db)

	defer db.Close(context.Background())

	book, err := repo.GetBookById(context.Background(), 2)
	if book != nil {
		t.Error("Expected to be nil for incorrect id.")
	}

	if err == nil {
		t.Error("Expected error for incorrect id.")
	}
}

func TestGetAllBooks(t *testing.T) {
	db, _ := NewConnectDb()
	repo := NewBookRepo(db)

	defer db.Close(context.Background())

	books, err := repo.GetAllBooks(context.Background())
	if len(books) == 0 {
		t.Fatal("Getted data is empty")
	}

	if err != nil {
		t.Fatalf("Error getted all books! %v", err)
	}
}

func TestUpdateBookById(t *testing.T) {
	db, _ := NewConnectDb()
	repo := NewBookRepo(db)

	defer db.Close(context.Background())

	book := domain.Book{
		Id:            1,
		Title:         "Новое название",
		Author:        "Новый автор",
		YearPublished: 2024,
	}

	err := repo.UpdateBookById(context.Background(), book)
	if err != nil {
		t.Fatalf("Error update book! %v", err)
	}
}

func TestDeleteBookById(t *testing.T) {
	db, _ := NewConnectDb()
	repo := NewBookRepo(db)

	defer db.Close(context.Background())

	err := repo.DeleteBookById(context.Background(), 1)
	if err != nil {
		t.Fatalf("Error delete book! %v", err)
	}
}

func TestDeleteAllBooks(t *testing.T) {
	db, _ := NewConnectDb()
	repo := NewBookRepo(db)

	defer db.Close(context.Background())

	err := repo.DeleteAllBooks(context.Background())
	if err != nil {
		t.Fatalf("Error delete all books! %v", err)
	}
}
