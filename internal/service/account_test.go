package service_test

import (
	"context"
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

func TestCreateAccount(t *testing.T) {

	dbPool := setupTestDB(t)
	defer func() {
		err := dbPool.Close()
		if err != nil {
			t.Errorf("failed to close DB connection: %v", err)
		}
	}()

	s := service.New(dbPool)
	ctx := context.Background()

	t.Run("create new account with valid details", func(t *testing.T) {

		accountName := "Test Savings Account"
		accountType := "bank"
		initialBalance := 1000.00

		account, err := s.CreateAccount(ctx, accountName, accountType, initialBalance)

		if err != nil {
			t.Fatalf("CreateAccount returned an error: %v", err)
		}

		if account.ID == "" {
			t.Error("expected valid account ID, got empty sting")
		}

		if account.LastUpdated.IsZero() {
			t.Error("expected valid last updated timestamp, got zero value")
		}

		if account.RegisteredOn.IsZero() {
			t.Error("expected valid registered on timestamp, got zero value")
		}

		AssertEqual(t, "invalid account name", accountName, account.Name)
		AssertEqual(t, "invalid account type", accountType, account.AccountType)
		AssertEqual(t, "invalid initial balance", initialBalance, account.CurrentBalance)
	})
}
