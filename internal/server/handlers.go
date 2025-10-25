package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IPOMPONI/go-base-CRUD/internal/bookstorage"
	"github.com/jackc/pgx/v5"
)

var dbConn *pgx.Conn

func NewHandler(conn *pgx.Conn) http.Handler {
	dbConn = conn

	mux := http.NewServeMux()

	mux.HandleFunc("POST /books", InsertBookHandler)
	mux.HandleFunc("GET /books", GetAllBooksHandler)
	mux.HandleFunc("GET /books/{id}", GetBookByIdHandler)
	mux.HandleFunc("DELETE /books", DeleteAllBooksHandler)
	mux.HandleFunc("DELETE /books/{id}", DeleteBookByIdHandler)
	return mux
}

func InsertBookHandler(w http.ResponseWriter, r *http.Request) {
	var book bookstorage.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := bookstorage.InsertBook(dbConn, book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message": "Book created"}`)
}

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := bookstorage.GetAllBooks(dbConn)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := bookstorage.GetBookById(dbConn, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func DeleteAllBooksHandler(w http.ResponseWriter, r *http.Request) {

	err := bookstorage.DeleteAllBooks(dbConn)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"message": "All books deleted"}`)
}

func DeleteBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = bookstorage.DeleteBookById(dbConn, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, `{"message": "Book with id = %v was deleted"}`, id)
}
