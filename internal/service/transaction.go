package service

import (
	"context"

	"github.com/VedantPatil1/gb-go/internal/db"
)

func (s *Service) GetTransaction(ctx context.Context, id string) (db.GetTransactionRow, error) {
	transaction, _ := s.queries.GetTransaction(ctx, id)

	return transaction, nil
}
