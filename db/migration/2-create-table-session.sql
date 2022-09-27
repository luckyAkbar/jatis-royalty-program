-- +migrate Up notransaction

CREATE TABLE IF NOT EXISTS "sessions" (
    id BIGINT PRIMARY KEY,
    access_token TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    expired_at TIMESTAMP NOT NULL

);

-- +migrate Down

DROP TABLE IF EXISTS "sessions";