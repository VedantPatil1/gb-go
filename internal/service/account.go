package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/VedantPatil1/gb-go/internal/db"
)

var (
	ErrAccountNotFound = errors.New("account not found")
)

func (s *Service) GetAccount(ctx context.Context, id string) (db.Account, error) {
	account, err := s.queries.GetAccount(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return db.Account{}, fmt.Errorf("account %s: %w", id, ErrAccountNotFound)
		}

		return db.Account{}, fmt.Errorf("failed to get account: %w", err)
	}

	return account, nil
}
