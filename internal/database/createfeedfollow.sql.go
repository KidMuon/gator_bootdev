// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: createfeedfollow.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
WITH insertedfeedfollow AS (
	INSERT INTO feedfollows (ID, created_at, updated_at, user_id, feed_id)
	VALUES (
		$1, 
		$2, 
		$3, 
		$4, 
		$5
	) RETURNING id, created_at, updated_at, user_id, feed_id
)
SELECT a.id, a.created_at, a.updated_at, a.user_id, a.feed_id, B.Name as UserName, C.Name as FeedName
FROM insertedfeedfollow A
INNER JOIN users B
	ON A.user_id = B.ID
INNER JOIN feeds C
	ON A.feed_id = C.ID
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

type CreateFeedFollowRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
	Username  string
	Feedname  string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
		&i.Username,
		&i.Feedname,
	)
	return i, err
}
