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

-- name: ListAccounts :many
SELECT id, external_id, name, account_type, current_balance, last_updated, registered_on
FROM accounts
ORDER BY name;
