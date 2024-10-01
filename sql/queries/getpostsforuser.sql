-- name: GetPostsForUser :many
SELECT A.*
FROM posts A
INNER JOIN feedfollows B
	ON A.feed_id = B.feed_id
WHERE B.user_id = $1
ORDER BY A.updated_at DESC
LIMIT $2;


