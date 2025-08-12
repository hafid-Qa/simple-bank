CREATE TABLE users (
    username VARCHAR NOT NULL PRIMARY KEY,
    hashed_password VARCHAR NOT NULL,
    full_name VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password_changed_at timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);


ALTER TABLE accounts ADD FOREIGN KEY (owner) REFERENCES users(username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
ALTER TABLE accounts ADD CONSTRAINT owner_currency_key UNIQUE (owner, currency);
