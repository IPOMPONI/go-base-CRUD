package main

import (
	"context"
	"fmt"

	"github.com/IPOMPONI/go-base-CRUD/internal/bookstorage"
)

func main() {
	db, err := bookstorage.NewConnectDb()

	if err != nil {
		println("Error connection: ", err)
		panic(err)
	}

	println("Connect!")

	err = bookstorage.InsertBook(db, bookstorage.Book{
		Title:         "Мастер и Маргарита",
		Author:        "Михаил Булгаков",
		YearPublished: 1966,
	})

	if err != nil {
		println("Error insert book_1")
	}

	err = bookstorage.InsertBook(db, bookstorage.Book{
		Title:         "Преступление и наказание",
		Author:        "Федор Достоевский",
		YearPublished: 1866,
	})

	if err != nil {
		println("Error insert book_2")
	}

	books, err := bookstorage.GetAllBooksData(db)

	fmt.Println(books)

	defer db.Close(context.Background())
}
