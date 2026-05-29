package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/VedantPatil1/gb-go/internal/service"
)

func TestGetAccount(t *testing.T) {

	dbPool := setupTestDB(t)
	defer func() {
		err := dbPool.Close()
		if err != nil {
			t.Errorf("failed to close DB connection: %v", err)
		}
	}()

	s := service.New(dbPool)
	ctx := context.Background()

	targetID := "test-account-id"

	_, err := dbPool.ExecContext(ctx, `
		INSERT INTO accounts (id, name, account_type, current_balance)
		VALUES (?, 'HDFC Salary Wallet', 'bank', 50000.00)
	`, targetID)

	t.Run("GetAccount returns correct account for given ID", func(t *testing.T) {
		if err != nil {
			t.Fatalf("failed to seed mock account data: %v", err)
		}
		account, err := s.GetAccount(ctx, targetID)

		if err != nil {
			t.Fatalf("GetAccount returned an error: %v", err)
		}

		if account.ID != targetID {
			t.Errorf("expected account ID %s, got %s", targetID, account.ID)
		}

		if account.Name != "HDFC Salary Wallet" {
			t.Errorf("expected account name 'HDFC Salary Wallet', got '%s'", account.Name)
		}

		if account.AccountType != "bank" {
			t.Errorf("expected account type 'bank', got '%s'", account.AccountType)
		}

		if account.CurrentBalance != 50000.00 {
			t.Errorf("expected current balance 50000.00, got %f", account.CurrentBalance)
		}

	})

	t.Run("GetAccount return appropriate error for non-existent account ID", func(t *testing.T) {
		_, err := s.GetAccount(ctx, "non-existent-id")

		if err == nil {
			t.Fatal("expected an error for non-existent account ID, got nil")
		}

		if !errors.Is(err, service.ErrAccountNotFound) {
			t.Errorf("expected error to be ErrAccountNotFound, got %v", err)
		}

	})
}
