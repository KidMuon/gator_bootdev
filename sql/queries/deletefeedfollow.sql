-- name: DeleteFeedFollow :one
DELETE FROM feedfollows
WHERE user_id = $1 AND feed_id = $2
RETURNING *;
