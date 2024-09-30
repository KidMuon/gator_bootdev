-- name: CreateFeedFollow :one
WITH insertedfeedfollow AS (
	INSERT INTO feedfollows (ID, created_at, updated_at, user_id, feed_url)
	VALUES (
		$1, 
		$2, 
		$3, 
		$4, 
		$5
	) RETURNING *
)
SELECT A.*, B.Name as UserName, C.Name as FeedName
FROM insertedfeedfollow A
INNER JOIN users B
	ON A.user_id = B.ID
INNER JOIN feeds C
	ON A.feed_url = C.url;
