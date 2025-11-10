-- +goose Up
-- +goose StatementBegin

-- Modify feeds table, these should be text as they can be longer than VARCHAR(255)
ALTER TABLE feeds 
    ALTER COLUMN url TYPE TEXT,
    ALTER COLUMN name TYPE TEXT;

-- Create the posts table
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT,
    published_at TIMESTAMP,
    feed_id UUID NOT NULL,
    CONSTRAINT fk_feed_id FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- Index for querying by feed and time
CREATE INDEX idx_posts_feed_published ON posts(feed_id, published_at DESC);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE feeds 
    ALTER COLUMN url TYPE VARCHAR(255),
    ALTER COLUMN name TYPE VARCHAR(255);

DROP TABLE posts;

-- +goose StatementEnd
