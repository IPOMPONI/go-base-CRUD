package main

import (
	"context"
	"log"
	"time"
	"net/http"
	"os"

	"booklib/internal/handler"
	"booklib/internal/middleware"
	"booklib/internal/repository/postgresql"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	db, err := postgresql.NewConnectDb(ctx)

	if err != nil {
		log.Println("Database connection failed:", err)
		return
	}

	log.Println("Database connected!")

	defer db.Close(ctx)

	bookRepo := postgresql.NewBookRepo(db)

	bookHandler := handler.NewBookHandler(bookRepo)

	mux := http.NewServeMux()

	bookHandler.InitRoutes(mux)

	handler := middleware.Chain(
		mux,
		middleware.RecoveryMiddleware,
		middleware.LoggingMiddleware,
	)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Port is not set. Default port: %s", port)
	}

	log.Println("Server starting on :" + port)
	log.Println(http.ListenAndServe(":" + port, handler))
}
