package main

import (
	"context"
	"log"
	"net/http"

	"booklib/internal/repository"
	"booklib/internal/server"
)

func main() {
	db, err := repository.NewConnectDb()

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Database connected!")

	defer db.Close(context.Background())

	bookRepo := repository.NewBookRepo(db)

	bookHandler := server.NewBookHandler(bookRepo)

	mux := http.NewServeMux()

	bookHandler.InitRoutes(mux)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
