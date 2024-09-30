-- +goose Up 
CREATE TABLE feedfollows (
	ID uuid PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	user_id uuid NOT NULL,
	feed_url VARCHAR(1023) NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY(feed_url) REFERENCES feeds(url) ON DELETE CASCADE,
	UNIQUE(user_id, feed_url)
);

-- +goose Down
DROP TABLE feedfollows;
