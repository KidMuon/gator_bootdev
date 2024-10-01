-- +goose Up
CREATE TABLE posts (
	ID uuid PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	title varchar(1024) NOT NULL,
	url varchar(1024) NULL,
	description varchar(1024) NULL,
	published_at TIMESTAMP NULL,
	feed_id uuid NOT NULL,
	UNIQUE(url),
	FOREIGN KEY(feed_id) REFERENCES feeds(ID)
);


-- +goose Down
DROP TABLE posts;
