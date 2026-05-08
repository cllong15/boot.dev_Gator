-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1, --id
    $2, --created_at
    $3, --updated_at
    $4, --title
	$5, --url
	$6, --description
	$7, --published_at
	$8  --feed_id
) ON CONFLICT (url) DO NOTHING
RETURNING *;
--

-- name: GetPostsForUser :many
SELECT posts.*
FROM posts
INNER JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = $1
ORDER BY published_at
LIMIT $2;
--
