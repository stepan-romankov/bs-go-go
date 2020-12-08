package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"testing"
)

var TestDbPool *pgxpool.Pool

func TestMain(m *testing.M) {
	dbUrl := os.Getenv("POSTGRES_URL")
	if dbUrl == "" {
		dbUrl = "postgres://postgres:test@localhost/blocksize_test?sslmode=disable&pool_max_conns=10"
	}
	dbPool, err := initDbPool(dbUrl)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	TestDbPool = dbPool
	defer dbPool.Close()

	err = Migrate(dbPool)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to execute migrations: %v\n", err)
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}

func CleanUpDb() {
	_, err := TestDbPool.Exec(context.Background(), "DELETE FROM "+SqlTable)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to cleanup DB: %v\n", err)
		os.Exit(1)
	}
}
