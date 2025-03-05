-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    name varchar(40) NOT NULL,
    url varchar(240) NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id)  ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds; 