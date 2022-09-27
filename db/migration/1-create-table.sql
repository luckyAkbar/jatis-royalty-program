-- +migrate Up notransaction

CREATE TABLE IF NOT EXISTS "users" (
    id BIGINT PRIMARY KEY,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- +migrate Down

DROP TABLE IF EXISTS "users";
