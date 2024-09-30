package main

import (
	"context"
	"fmt"

	"github.com/kidmuon/gator_bootdev/internal/database"
)

func handlerUnfollow(state *State, command Command, user database.User) error {
	if len(command.args) < 1 {
		return fmt.Errorf("unfollow expects a url to unfollow")
	}

	feedurl := command.args[0]
	feed, err := state.db.GetFeedByURL(context.Background(), feedurl)
	if err != nil {
		return fmt.Errorf("Unable to find feed with url: \"%s\", error: %v", feedurl, err)
	}

	feedfollowToDelete := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	_, err = state.db.DeleteFeedFollow(context.Background(), feedfollowToDelete)
	if err != nil {
		return fmt.Errorf("problem unfollowing %s: %v", feedurl, err)
	}

	fmt.Printf("User %s no longer follows %s", user.Name, feedurl)

	return nil
}
