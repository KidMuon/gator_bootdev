-- name: GetFeedFollowsForUser :many
SELECT C.name feedname, B.name username
FROM feedfollows A
INNER JOIN users B
	ON A.user_id = B.ID
INNER JOIN feeds C
	ON A.feed_id = C.ID
WHERE B.name = $1;
