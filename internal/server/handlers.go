package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IPOMPONI/go-base-CRUD/internal/bookstorage"
	"github.com/jackc/pgx/v5"
)

var dbConn *pgx.Conn

func NewHandler(conn *pgx.Conn) http.Handler {
	dbConn = conn

	mux := http.NewServeMux()

	mux.HandleFunc("POST /books", InsertBookHandler)
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
