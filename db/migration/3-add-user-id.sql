-- +migrate Up notransaction
ALTER TABLE "sessions" ADD COLUMN IF NOT EXISTS user_id BIGINT NOT NULL;

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");


-- +migrate Down
ALTER TABLE "sessions" DROP COLUMN IF EXISTS "user_id";

