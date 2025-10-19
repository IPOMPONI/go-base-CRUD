package bookstorage

import (
	"context"
	"testing"
)

func TestInsertBook(t *testing.T) {
	db, _ := NewConnectDb()

	_, err := db.Exec(context.Background(), "TRUNCATE TABLE books RESTART IDENTITY")

	if err != nil {
		t.Fatalf("Error truncate table!")
	}

	defer db.Close(context.Background())

	book := Book{
		Title:         "Книга 1",
		Author:        "Автор 1",
		YearPublished: 2025,
	}

	err = InsertBook(db, book)

	if err != nil {
		t.Fatalf("Error insert book! %v", err)
	}
}

func TestInsertBookExistingTitle(t *testing.T) {
	db, _ := NewConnectDb()

	defer db.Close(context.Background())

	book := Book{
		Title:         "Книга 1",
		Author:        "Автор 1",
		YearPublished: 2025,
	}

	err := InsertBook(db, book)

	if err == nil {
		t.Error("Expected error for existing title")
	}
}

func TestInsertBookZeroYear(t *testing.T) {
	db, _ := NewConnectDb()

	defer db.Close(context.Background())

	book := Book{
		Title:         "Книга 1",
		Author:        "Автор 1",
		YearPublished: 0,
	}

	err := InsertBook(db, book)

	if err == nil {
		t.Error("Expected error for zero year.")
	}
}

func TestInsertBookFutureYear(t *testing.T) {
	db, _ := NewConnectDb()

	defer db.Close(context.Background())

	book := Book{
		Title:         "Книга 1",
		Author:        "Автор 1",
		YearPublished: 2077,
	}

	err := InsertBook(db, book)

	if err == nil {
		t.Error("Expected error for future year.")
	}
}

func TestGetBookById(t *testing.T) {
	db, _ := NewConnectDb()

	defer db.Close(context.Background())

	var book *Book

	book, err := GetBookById(db, 1)

	if book == nil {
		t.Fatalf("Geted book is nil!")
	}

	if err != nil {
		t.Fatalf("Error get book by id! %v", err)
	}
}

func TestGetBookByIncorrectId(t *testing.T) {
	db, _ := NewConnectDb()

	defer db.Close(context.Background())

	var book *Book

	book, err := GetBookById(db, 2)

	if book != nil {
		t.Error("Expected to be nil for incorrect id.")
	}

	if err == nil {
		t.Error("Expected error for incorrect id.")
	}
}
