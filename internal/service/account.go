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

func (s *Service) CreateAccount(ctx context.Context, name string, accountType string, initialBalance float64) (db.Account, error) {

	accountPayload := db.CreateAccountParams{
		ID:             "test-id",
		Name:           name,
		AccountType:    accountType,
		CurrentBalance: initialBalance,
	}
	account, _ := s.queries.CreateAccount(ctx, accountPayload)

	return account, nil

}
