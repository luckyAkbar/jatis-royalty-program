-- +migrate Up notransaction

CREATE TABLE IF NOT EXISTS "vouchers" (
    id BIGINT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    invoice_id BIGINT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    expired_at TIMESTAMP NOT NULL,
    is_claimed BOOLEAN DEFAULT FALSE,
    value BIGINT NOT NULL
);

ALTER TABLE "vouchers" ADD FOREIGN KEY (user_id) REFERENCES "users" ("id");
ALTER TABLE "vouchers" ADD FOREIGN KEY (invoice_id) REFERENCES "invoices" ("id");

-- +migrate Down

DROP TABLE IF EXISTS "vouchers";