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

		AssertEqual(t, "invalid account ID", targetID, account.ID)
		AssertEqual(t, "invalid account name", "HDFC Salary Wallet", account.Name)
		AssertEqual(t, "invalid account type", "bank", account.AccountType)
		AssertEqual(t, "invalid current balance", 50000.00, account.CurrentBalance)
	})

	t.Run("GetAccount return appropriate error for non-existent account ID", func(t *testing.T) {
		_, err := s.GetAccount(ctx, "non-existent-id")

		AssertError(t, service.ErrAccountNotFound, err)
	})
}

func AssertEqual(t *testing.T, msg string, want, got any) {
	t.Helper()

	if want != got {
		t.Errorf("%s: expected %v, got %v", msg, want, got)
	}
}

func AssertError(t *testing.T, want, got error)  {
	t.Helper()

	if got == nil {
		t.Errorf("no error returned: expected %v, got %v", want, got)
	}

	if !errors.Is(got, want) {
		t.Errorf("unexpected error type: expected %v, got %v", want, got)
	}
}
