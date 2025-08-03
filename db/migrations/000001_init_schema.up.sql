
-- Create the initial schema for the accounting system
CREATE TABLE accounts (
    id bigserial PRIMARY KEY,
    owner VARCHAR NOT NULL,
    balance BIGINT NOT NULL DEFAULT 0,
    currency VARCHAR NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE entries (
    id bigserial PRIMARY KEY,
    account_id BIGINT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    amount BIGINT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE transfers (
    id bigserial PRIMARY KEY,
    from_account_id BIGINT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    to_account_id BIGINT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    amount BIGINT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE entries ADD FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE;
ALTER TABLE transfers ADD FOREIGN KEY (from_account_id) REFERENCES accounts(id) ON DELETE CASCADE;
ALTER TABLE transfers ADD FOREIGN KEY (to_account_id) REFERENCES accounts(id) ON DELETE CASCADE;

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_accounts_owner ON accounts(owner);
CREATE INDEX IF NOT EXISTS idx_entries_account_id ON entries(account_id);
CREATE INDEX IF NOT EXISTS idx_transfers_from_account_id ON transfers(from_account_id);
CREATE INDEX IF NOT EXISTS idx_transfers_to_account_id ON transfers(to_account_id);
CREATE INDEX IF NOT EXISTS idx_transfers_from_to ON transfers(from_account_id, to_account_id);

COMMENT ON COLUMN entries.amount IS 'The amount of the entry, can be positive or negative';
COMMENT ON COLUMN transfers.amount IS 'The amount of the transfer, should be positive';
