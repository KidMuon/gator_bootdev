package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kidmuon/gator_bootdev/internal/database"
	"time"
)

func handlerFollow(state *State, command Command, user database.User) error {
	if len(command.args) < 1 {
		return fmt.Errorf("follow requires a url for the user to follow")
	}

	feedurl := command.args[0]

	feedFollowToCreate := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedUrl:   feedurl,
	}

	feedFollowRow, err := state.db.CreateFeedFollow(context.Background(), feedFollowToCreate)
	if err != nil {
		return fmt.Errorf("error creating feedfollow: %v", err)
	}

	fmt.Printf("Feed \"%s\" now followed by %s \n", feedFollowRow.Feedname, feedFollowRow.Username)
	return nil
}
