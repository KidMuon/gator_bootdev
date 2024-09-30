package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kidmuon/gator_bootdev/internal/database"
)

func handlerFollow(state *State, command Command, user database.User) error {
	if len(command.args) < 1 {
		return fmt.Errorf("follow requires a url for the user to follow")
	}

	feedurl := command.args[0]
	feed, err := state.db.GetFeedByURL(context.Background(), feedurl)
	if err != nil {
		return fmt.Errorf("Unable to find feed with url: \"%s\", error: %v", feedurl, err)
	}

	feedFollowToCreate := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollowRow, err := state.db.CreateFeedFollow(context.Background(), feedFollowToCreate)
	if err != nil {
		return fmt.Errorf("error creating feedfollow: %v", err)
	}

	fmt.Printf("Feed \"%s\" now followed by %s \n", feedFollowRow.Feedname, feedFollowRow.Username)
	return nil
}
