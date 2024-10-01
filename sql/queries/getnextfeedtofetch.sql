-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY updated_at NULLS FIRST
LIMIT 1;
