-- +goose Up
CREATE TABLE feeds (
	name VARCHAR(255) NOT NULL,
	url VARCHAR(1023) UNIQUE NOT NULL,
	user_id uuid NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds


