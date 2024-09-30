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

	feedUrl := command.args[0]

	feedfollowToDelete := database.DeleteFeedFollowParams{
		UserID:  user.ID,
		FeedUrl: feedUrl,
	}

	ff, err := state.db.DeleteFeedFollow(context.Background(), feedfollowToDelete)
	if err != nil {
		return fmt.Errorf("problem unfollowing %s: %v", feedUrl, err)
	}

	fmt.Printf("User %s no longer follows %s", user.Name, ff.FeedUrl)

	return nil
}
