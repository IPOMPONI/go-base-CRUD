package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"booklib/internal/handler"
	"booklib/internal/repository/postgresql"
)

func main() {
	db, err := postgresql.NewConnectDb()

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Database connected!")

	defer db.Close(context.Background())

	bookRepo := postgresql.NewBookRepo(db)

	bookHandler := handler.NewBookHandler(bookRepo)

	mux := http.NewServeMux()

	bookHandler.InitRoutes(mux)

	log.Println("Server starting on :" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}
