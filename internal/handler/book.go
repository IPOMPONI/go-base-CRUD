package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"booklib/internal/domain"
)

type BookHandler struct {
	repo domain.BookRepo
}

func NewBookHandler(repo domain.BookRepo) *BookHandler {
	return &BookHandler{repo: repo}
}

func (bh *BookHandler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /books", bh.insertBook)
	mux.HandleFunc("GET /books/{id}", bh.getBookById)
	mux.HandleFunc("GET /books", bh.getAllBooks)
	mux.HandleFunc("PUT /books/{id}", bh.updateBookById)
	mux.HandleFunc("DELETE /books/{id}", bh.deleteBookById)
	mux.HandleFunc("DELETE /books", bh.deleteAllBooks)
}

func (bh *BookHandler) insertBook(w http.ResponseWriter, r *http.Request) {
	var book domain.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := bh.repo.InsertBook(r.Context(), book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message": "Book created"}`)
}

func (bh *BookHandler) getBookById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := bh.repo.GetBookById(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (bh *BookHandler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := bh.repo.GetAllBooks(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (bh *BookHandler) updateBookById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var book domain.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	book.Id = id

	err = bh.repo.UpdateBookById(r.Context(), book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Book with id = %d was updated"}`, id)
}

func (bh *BookHandler) deleteBookById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = bh.repo.DeleteBookById(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, `{"message": "Book with id = %v was deleted"}`, id)
}

func (bh *BookHandler) deleteAllBooks(w http.ResponseWriter, r *http.Request) {

	err := bh.repo.DeleteAllBooks(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"message": "All books deleted"}`)
}
