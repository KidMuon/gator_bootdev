// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: deletefeedfollow.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const deleteFeedFollow = `-- name: DeleteFeedFollow :one
DELETE FROM feedfollows
WHERE User_ID = $1 AND Feed_Url = $2
RETURNING id, created_at, updated_at, user_id, feed_url
`

type DeleteFeedFollowParams struct {
	UserID  uuid.UUID
	FeedUrl string
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) (Feedfollow, error) {
	row := q.db.QueryRowContext(ctx, deleteFeedFollow, arg.UserID, arg.FeedUrl)
	var i Feedfollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedUrl,
	)
	return i, err
}
