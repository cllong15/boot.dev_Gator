-- name: UpdateFeed :exec
UPDATE feeds
SET updated_at = $1
WHERE id = $2;
UPDATE feeds
SET last_fetched_at = $1
WHERE feed_id = $2;
UPDATE feed_follows
SET updated_at = $1
WHERE feed_id = $2;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds ORDER BY last_fetched_at NULLS FIRST;