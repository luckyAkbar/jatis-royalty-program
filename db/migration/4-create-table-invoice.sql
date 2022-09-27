-- +migrate Up notransaction

CREATE TABLE IF NOT EXISTS "invoices" (
    id BIGINT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    total_price BIGINT NOT NULL
);

ALTER TABLE "invoices" ADD FOREIGN KEY (user_id) REFERENCES "users" (id);

-- +migrate Down

DROP TABLE IF EXISTS "invoices";