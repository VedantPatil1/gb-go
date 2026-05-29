package service_test

import (
	"database/sql"
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

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
