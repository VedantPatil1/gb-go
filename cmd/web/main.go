package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/VedantPatil1/gb-go/internal/server"
	_ "modernc.org/sqlite"
)

const dbPath = "../../gb-go.db"
const schemaPath = "db/schema.sql"

func openDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dns)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()

		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil

}

func migrate(ctx context.Context, db *sql.DB) error {
	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("read schema: %w", err)
	}

	if _, err := db.Exec(string(schemaBytes)); err != nil {
		return fmt.Errorf("excecute schema: %w", err)
	}

	return nil
}

func main() {
	ctx := context.Background()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := openDB(dbPath)
	if err != nil {
		logger.Error("failed to open database", "error", err)
		os.Exit(1)
	}

	logger.Info("connected to database")
	defer db.Close()

	if err := migrate(ctx, db); err != nil {
		logger.Error("database migration failed", "error", err)
		os.Exit(1)
	}
	logger.Info("database migration done successfully")

	server := server.NewServer(
		server.Config{Addr: ":8001"},
		server.Dependencies{Logger: logger, Db: db})
	logger.Info("starting server", "addr", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		logger.Error("server stopped", "error", err)
	}
}
