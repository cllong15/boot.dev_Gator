-- +goose Up
CREATE TABLE feeds (
	id UUID primary key,
	created_at timestamp not null,
	updated_at timestamp not null,
	name text not null,
	url text unique not null,
	user_id UUID not null references users
		on delete cascade
		on update cascade,
	foreign key (user_id) references users (id)

);

-- +goose Down
DROP TABLE feeds;
