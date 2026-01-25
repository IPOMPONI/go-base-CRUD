package handler

import (
	"encoding/json"
	"net/http"

	"booklib/internal/domain"
	"booklib/internal/middleware"
)

type BookHandler struct {
	repo domain.BookRepo
}

func NewBookHandler(repo domain.BookRepo) *BookHandler {
	return &BookHandler{repo: repo}
}

func (bh *BookHandler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /books", bh.insertBook)
	mux.HandleFunc("GET /books", bh.getAllBooks)
	mux.HandleFunc("DELETE /books", bh.deleteAllBooks)
	mux.Handle("GET /books/{id}", middleware.CheckBookIdMiddleware(http.HandlerFunc(bh.getBookById)))
	mux.Handle("PUT /books/{id}", middleware.CheckBookIdMiddleware(http.HandlerFunc(bh.updateBookById)))
	mux.Handle("DELETE /books/{id}", middleware.CheckBookIdMiddleware(http.HandlerFunc(bh.deleteBookById)))
}

func (bh *BookHandler) insertBook(w http.ResponseWriter, r *http.Request) {
	var book domain.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON!", http.StatusBadRequest)
		return
	}

	if err := bh.repo.InsertBook(r.Context(), book); err != nil {
		http.Error(w, "Data insertion error!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (bh *BookHandler) getBookById(w http.ResponseWriter, r *http.Request) {
 	id, ok := r.Context().Value("bookId").(int)

  	if !ok {
        http.Error(w, "Book 'id' not found in context!", http.StatusInternalServerError)
        return
    }

	book, err := bh.repo.GetBookById(r.Context(), id)

	if err != nil {
		http.Error(w, "Error getting the book by the specified 'id'!", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (bh *BookHandler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := bh.repo.GetAllBooks(r.Context())

	if err != nil {
		http.Error(w, "Error getting the books!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (bh *BookHandler) updateBookById(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("bookId").(int)

	if !ok {
        http.Error(w, "Book 'id' not found in context!", http.StatusInternalServerError)
        return
    }

	var book domain.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON!", http.StatusBadRequest)
		return
	}

	book.Id = id

	err := bh.repo.UpdateBookById(r.Context(), book)

	if err != nil {
		http.Error(w, "Error updating the book by the specified 'id'!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (bh *BookHandler) deleteBookById(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("bookId").(int)

	if !ok {
        http.Error(w, "Book 'id' not found in context!", http.StatusInternalServerError)
        return
    }

	err := bh.repo.DeleteBookById(r.Context(), id)

	if err != nil {
		http.Error(w, "Error deleting the book by the specified 'id'!", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (bh *BookHandler) deleteAllBooks(w http.ResponseWriter, r *http.Request) {

	err := bh.repo.DeleteAllBooks(r.Context())

	if err != nil {
		http.Error(w, "Error deleting the books!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
