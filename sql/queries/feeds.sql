-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeedByName :one
SELECT * FROM feeds
WHERE name = $1;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeedByID :one
SELECT * FROM feeds
WHERE id = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsWithUsers :many
SELECT 
    feeds.*,
    users.name AS user_name
FROM feeds
INNER JOIN users ON feeds.user_id = users.id;

-- name: ResetFeeds :execrows
DELETE FROM feeds;


-- Updated current feed as fetched at. 
-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $1, updated_at = $2
WHERE id = $3;

-- Should return the next feed we should fetch posts from. 
-- We want to scrape all the feeds in a continuous loop. 
-- A simple approach is to keep track of when a feed was 
-- last fetched, and always fetch the oldest one first 
-- (or any that haven't ever been fetched). 

-- name: GetNextFeedToFetch :one
SELECT * 
FROM feeds
ORDER BY last_fetched_at DESC NULLS FIRST;