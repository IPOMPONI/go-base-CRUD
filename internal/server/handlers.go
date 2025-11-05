package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"booklib/internal/model"
	"booklib/internal/repository"
)

type BookHandler struct {
	repo repository.BookRepo
}

func NewBookHandler(repo repository.BookRepo) *BookHandler {
	return &BookHandler{repo: repo}
}

func (bh *BookHandler) NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /books", bh.InsertBookHandler)
	mux.HandleFunc("GET /books/{id}", bh.GetBookByIdHandler)
	mux.HandleFunc("GET /books", bh.GetAllBooksHandler)
	mux.HandleFunc("PUT /books/{id}", bh.UpdateBookByIdHandler)
	mux.HandleFunc("DELETE /books/{id}", bh.DeleteBookByIdHandler)
	mux.HandleFunc("DELETE /books", bh.DeleteAllBooksHandler)
	return mux
}

func (bh *BookHandler) InsertBookHandler(w http.ResponseWriter, r *http.Request) {
	var book model.Book

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

func (bh *BookHandler) GetBookByIdHandler(w http.ResponseWriter, r *http.Request) {
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

func (bh *BookHandler) GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := bh.repo.GetAllBooks(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (bh *BookHandler) UpdateBookByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	err = bh.repo.UpdateBookById(r.Context(), book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Book with id = %d was updated"}`, id)
}

func (bh *BookHandler) DeleteBookByIdHandler(w http.ResponseWriter, r *http.Request) {
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

func (bh *BookHandler) DeleteAllBooksHandler(w http.ResponseWriter, r *http.Request) {

	err := bh.repo.DeleteAllBooks(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"message": "All books deleted"}`)
}
