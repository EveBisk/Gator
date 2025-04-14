-- +goose Up
ALTER TABLE feeds
ADD COLUMN IF NOT EXISTS last_fetched_at timestamp NULL;

-- +goose Down
ALTER TABLE feeds
DROP COLUMN IF EXISTS last_fetched_at;