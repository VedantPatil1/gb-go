-- name: GetAccount :one
SELECT id, external_id, name, account_type, current_balance, last_updated, registered_on
FROM accounts
WHERE id = ?;

-- name: CreateAccount :one
INSERT INTO accounts (id, external_id, name, account_type, current_balance)
VALUES (?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateAccount :one
UPDATE accounts
SET name = ?,
    account_type = ?,
    last_updated = CURRENT_TIMESTAMP
WHERE id = ? RETURNING *;

-- name: UpdateCurrentBalance :one
UPDATE accounts
SET current_balance = ?,
    last_updated = CURRENT_TIMESTAMP
WHERE id = ? RETURNING *;

-- name: CreateTransaction :one
INSERT INTO transactions (id, account_id, amount, description, metadata, tags, is_offset, date_of_transaction)
VALUES (?, ?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: GetTransaction :one
SELECT id, account_id, amount, description, metadata, tags, is_offset, date_of_transaction, updated_at
FROM transactions
WHERE id = ?;

-- name: UpdateTransaction :one
UPDATE transactions
SET description = ?,
    metadata = ?,
    tags = ?,
    is_offset = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ? RETURNING *;

-- name: ListTransactionsByAccount :many
SELECT id, account_id, amount, description, metadata, tags, is_offset, date_of_transaction, updated_at
FROM transactions
WHERE account_id = ?
ORDER BY date_of_transaction DESC
LIMIT ?;
