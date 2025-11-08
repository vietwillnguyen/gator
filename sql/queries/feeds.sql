-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE name = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: ResetFeeds :execrows
DELETE FROM feeds;

-- name: GetFeedsWithUsers :many
SELECT 
    feeds.*,
    users.name AS user_name
FROM feeds
INNER JOIN users ON feeds.user_id = users.id;