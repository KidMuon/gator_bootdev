-- name: MarkFeedFetched :one
UPDATE feeds
SET updated_at = $1,
 last_fetched_at = $2
WHERE ID = $3
RETURNING *;
