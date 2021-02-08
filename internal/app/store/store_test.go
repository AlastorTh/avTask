package store_test

import (
	"os"
	"testing"
)

var (
	connectString string
)

func TestMain(m *testing.M) {
	connectString = os.Getenv("DATABASE_URL")
	if connectString == "" {
		connectString = "host=localhost user=postgres password=pass port=5432 dbname=postgres_test sslmode=disable"
	}

	os.Exit(m.Run())
}
