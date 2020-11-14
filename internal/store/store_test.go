package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(t *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_test sslmode=disable user=postgres password=123"
	}

	os.Exit(t.Run())
}
