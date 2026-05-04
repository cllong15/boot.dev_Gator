-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id uuid NOT NULL references users ON DELETE CASCADE,
	feed_id uuid NOT NULL references feeds ON DELETE CASCADE,
	constraint user_feed
		UNIQUE (user_id, feed_id)

);

-- +goose Down
DROP TABLE feed_follows;
