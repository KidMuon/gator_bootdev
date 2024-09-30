-- name: ListFeeds :many
SELECT A.Name as Feed_Name, A.Url, B.Name as User_Name
FROM feeds A
INNER JOIN users B
	ON A.user_id = B.ID;
