// Package service orchestrates the core business logic of the application.
package service

import (
	"database/sql"

	"github.com/VedantPatil1/gb-go/internal/db"
)

type Service struct {
	db      *sql.DB
	queries *db.Queries
}

func New(dbPool *sql.DB) *Service {
	return &Service{
		db:      dbPool,
		queries: db.New(dbPool),
	}
}
