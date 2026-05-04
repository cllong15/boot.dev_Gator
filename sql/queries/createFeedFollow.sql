-- name: CreateFeedFollow :many
WITH inserted_feed_follows AS (
	INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
	VALUES (
	$1,
	$2,
	$3,
	$4,
	$5
	)
	RETURNING *
)
SELECT
    inserted_feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follows
INNER JOIN feeds on feeds.id = inserted_feed_follows.feed_id
INNER JOIN users on users.id = inserted_feed_follows.user_id;

-- It should return all the feed follows for a given user
-- and include the names of the feeds and user in the result.

-- name: GetFeedFollowsForUser :many
SELECT
	feed_follows.*,
	feeds.name as feed_name,
	users.name as user_name
FROM feed_follows
INNER JOIN feeds on feeds.id = feed_follows.feed_id
INNER JOIN users on users.id = feed_follows.user_id
WHERE users.name = $1;