-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    name varchar(40) NOT NULL,
    CONSTRAINT un_user_name UNIQUE(name)
);

-- +goose Down
DROP TABLE users; 