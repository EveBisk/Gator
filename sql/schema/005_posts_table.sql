-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    title TEXT NOT NULL,
    url varchar(240) NOT NULL,    
    description TEXT,
    published_at timestamp NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id)  ON DELETE CASCADE,
    CONSTRAINT un_post_url UNIQUE(url)
);

-- +goose Down
DROP TABLE posts; 