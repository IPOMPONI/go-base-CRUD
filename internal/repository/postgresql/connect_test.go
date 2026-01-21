package postgresql

import (
	"time"
	"context"
	"testing"
)

func TestConnectDb(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	db, err := NewConnectDb(ctx)

	if err != nil {
		t.Fatalf("Connection failed: %v", err)
	}

	defer db.Close(ctx)

	err = db.Ping(ctx)

	if err != nil {
		t.Fatalf("Failed ping db: %v", err)
	}
}
