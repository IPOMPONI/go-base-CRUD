package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"booklib/internal/bookstorage"
	"booklib/internal/model"

	"github.com/jackc/pgx/v5"
)

type HandlerData struct {
	DbConn *pgx.Conn
}

func (hD HandlerData) NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /books", hD.InsertBookHandler)
	mux.HandleFunc("GET /books/{id}", hD.GetBookByIdHandler)
	mux.HandleFunc("GET /books", hD.GetAllBooksHandler)
	mux.HandleFunc("PUT /books/{id}", hD.UpdateBookByIdHandler)
	mux.HandleFunc("DELETE /books/{id}", hD.DeleteBookByIdHandler)
	mux.HandleFunc("DELETE /books", hD.DeleteAllBooksHandler)
	return mux
}

func (hD HandlerData) InsertBookHandler(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := bookstorage.InsertBook(hD.DbConn, book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message": "Book created"}`)
}

func (hD HandlerData) GetBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := bookstorage.GetBookById(hD.DbConn, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (hD HandlerData) GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := bookstorage.GetAllBooks(hD.DbConn)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (hD HandlerData) UpdateBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var book model.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	book.Id = id

	err = bookstorage.UpdateBookById(hD.DbConn, book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Book with id = %d was updated"}`, id)
}

func (hD HandlerData) DeleteBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = bookstorage.DeleteBookById(hD.DbConn, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, `{"message": "Book with id = %v was deleted"}`, id)
}

func (hD HandlerData) DeleteAllBooksHandler(w http.ResponseWriter, r *http.Request) {

	err := bookstorage.DeleteAllBooks(hD.DbConn)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"message": "All books deleted"}`)
}
