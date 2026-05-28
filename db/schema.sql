CREATE TABLE IF NOT EXISTS accounts (
    id TEXT PRIMARY KEY,
    external_id TEXT UNIQUE,
    name TEXT NOT NULL UNIQUE,
    account_type TEXT NOT NULL,                   -- Treat purely as a TEXT primitive
    current_balance REAL NOT NULL DEFAULT 0.0,
    last_updated DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    registered_on DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
