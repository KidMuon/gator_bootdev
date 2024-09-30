package main

import (
	"context"
	"fmt"
)

func handlerFollowing(state *State, command Command) error {
	feedfollows, err := state.db.GetFeedFollowsForUser(context.Background(), state.cfg.Username)

	if err != nil {
		return fmt.Errorf("problem getting feed follows from database: %v", err)
	}

	for _, ff := range feedfollows {
		fmt.Printf("user: %s feed: \"%s\"\n", ff.Username, ff.Feedname)
	}

	return nil
}
