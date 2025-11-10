-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- Get posts for a user. 
-- feeds may have several posts
-- users can follow feeds
-- feed_follows tracks users and what feeds they follow
-- order by most recent, e.g. time value is highest

-- name: GetPostsForUser :many
SELECT posts.*
FROM posts
INNER JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
INNER JOIN users ON users.id = feed_follows.user_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC NULLS FIRST;


-- name: GetPostByURL :one
SELECT * 
FROM posts
WHERE url = $1;
