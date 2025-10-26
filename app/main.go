package main

import (
	"context"
	"log"
	"net/http"

	"github.com/IPOMPONI/go-base-CRUD/internal/bookstorage"
	"github.com/IPOMPONI/go-base-CRUD/internal/server"
)

func main() {
	db, err := bookstorage.NewConnectDb()

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Database connected!")

	defer db.Close(context.Background())

	handlerData := server.HandlerData{DbConn: db}

	handler := handlerData.NewHandler()

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
