// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: getfeedfollowforuser.sql

package database

import (
	"context"
)

const getFeedFollowsForUser = `-- name: GetFeedFollowsForUser :many
SELECT C.name feedname, B.name username
FROM feedfollows A
INNER JOIN users B
	ON A.user_id = B.ID
INNER JOIN feeds C
	ON A.feed_url = C.url
WHERE B.name = $1
`

type GetFeedFollowsForUserRow struct {
	Feedname string
	Username string
}

func (q *Queries) GetFeedFollowsForUser(ctx context.Context, name string) ([]GetFeedFollowsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUser, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedFollowsForUserRow
	for rows.Next() {
		var i GetFeedFollowsForUserRow
		if err := rows.Scan(&i.Feedname, &i.Username); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}