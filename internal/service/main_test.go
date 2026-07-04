package service_test

import (
	"database/sql"
	"errors"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

const schemaPath = "../../db/schema.sql"

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	dbPool, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		t.Fatalf("Failed to read schema file: %v", err)
	}

	if _, err := dbPool.Exec(string(schemaBytes)); err != nil {
		t.Fatalf("Failed to execute schema: %v", err)
	}

	return dbPool
}

func AssertEqual(t *testing.T, msg string, want, got any) {
	t.Helper()

	if want != got {
		t.Errorf("%s: expected %v, got %v", msg, want, got)
	}
}

func AssertError(t *testing.T, want, got error) {
	t.Helper()

	if got == nil {
		t.Errorf("no error returned: expected %v, got %v", want, got)
	}

	if !errors.Is(got, want) {
		t.Errorf("unexpected error type: expected %v, got %v", want, got)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
