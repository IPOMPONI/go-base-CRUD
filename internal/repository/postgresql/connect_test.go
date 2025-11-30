package postgresql

import (
	"context"
	"testing"
)

func TestConnectDb(t *testing.T) {
	db, err := NewConnectDb()

	if err != nil {
		t.Fatalf("Connection failed: %v", err)
	}

	defer db.Close(context.Background())

	err = db.Ping(context.Background())

	if err != nil {
		t.Fatalf("Failed ping db: %v", err)
	}
}
