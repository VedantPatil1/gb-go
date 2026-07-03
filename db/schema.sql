-- =============================
-- TABLE: accounts
-- Stores core information and current balance of an account.
-- =============================
CREATE TABLE IF NOT EXISTS accounts (
    id TEXT PRIMARY KEY,
    external_id TEXT UNIQUE,
    name TEXT NOT NULL UNIQUE,
    account_type TEXT NOT NULL,
    current_balance REAL NOT NULL DEFAULT 0.0,
    last_updated DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    registered_on DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- =============================
-- TABLE: transactions
-- Audit log of all financial movements and state changes.
-- Includes soft-delete capability for audit integrity.
-- =============================
CREATE TABLE IF NOT EXISTS transactions (
    id TEXT PRIMARY KEY,
    account_id TEXT NOT NULL,
    amount REAL NOT NULL,
    description TEXT,
    metadata JSON,
    tags TEXT,
    is_offset BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE, -- Soft delete flag (1=true)
    date_of_transaction DATETIME NOT NULL,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(account_id) REFERENCES accounts(id) ON DELETE RESTRICT
);

-- =============================
-- INDEXES: Performance Lookups
-- This section is strictly cleaned to contain no duplicates.
-- =============================
CREATE INDEX IF NOT EXISTS idx_transactions_by_account ON transactions (account_id);
CREATE INDEX IF NOT EXISTS idx_transactions_by_date ON transactions (date_of_transaction);