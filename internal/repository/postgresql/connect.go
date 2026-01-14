package postgresql

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewConnectDb() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), "user="+os.Getenv("USER")+" dbname="+os.Getenv("DBNAME")+" sslmode=disable")
}
