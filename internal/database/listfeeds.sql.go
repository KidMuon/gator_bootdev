// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: listfeeds.sql

package database

import (
	"context"
)

const listFeeds = `-- name: ListFeeds :many
SELECT A.Name as Feed_Name, A.Url, B.Name as User_Name
FROM feeds A
INNER JOIN users B
	ON A.user_id = B.ID
`

type ListFeedsRow struct {
	FeedName string
	Url      string
	UserName string
}

func (q *Queries) ListFeeds(ctx context.Context) ([]ListFeedsRow, error) {
	rows, err := q.db.QueryContext(ctx, listFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListFeedsRow
	for rows.Next() {
		var i ListFeedsRow
		if err := rows.Scan(&i.FeedName, &i.Url, &i.UserName); err != nil {
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
