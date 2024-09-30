-- name: DeleteFeedFollow :one
DELETE FROM feedfollows
WHERE User_ID = $1 AND Feed_Url = $2
RETURNING *;
