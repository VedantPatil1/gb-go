package service_test

import (
	"context"
	"testing"

	"github.com/VedantPatil1/gb-go/internal/service"
)

func TestGetTransaction(t *testing.T) {

	dbPool := setupTestDB(t)
	defer func() {
		err := dbPool.Close()
		if err != nil {
			t.Errorf("failed to close DB connection: %v", err)
		}
	}()

	s := service.New(dbPool)
	ctx := context.Background()

	testAccount, _ := s.CreateAccount(ctx, "test account", "bank", 100.00)

	transactionID := "test-transaction-id"

	query := `
	INSERT INTO transactions (id, account_id, amount, description, date_of_transaction)
	VALUES (?, ?, 50.00, 'Test transaction', '2023-10-27 10:00:00')
	`

	_, err := dbPool.ExecContext(ctx, query, transactionID, testAccount.ID)

	if err != nil {
		t.Fatalf("failed to seed mock transaction data: %v", err)
	}

	t.Run("GetTransaction returns correct transaction for given ID", func(t *testing.T) {
		transaction, err := s.GetTransaction(ctx, transactionID)

		if err != nil {
			t.Fatalf("GetTransaction returned an error: %v", err)
		}

		AssertEqual(t, "invalid transaction ID", transactionID, transaction.ID)

	})

}
