package postgresql

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewConnectDb(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "user="+os.Getenv("USER")+" dbname="+os.Getenv("DBNAME")+" sslmode=disable")
}
