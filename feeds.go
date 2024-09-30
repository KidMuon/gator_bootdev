package main

import (
	"context"
	"fmt"
)

func handlerFeeds(state *State, command Command) error {
	feeds, err := state.db.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed getting list of feeds from database: %v")
	}

	for _, feed := range feeds {
		fmt.Printf("Feed: %s, Url: %s, Username: %s\n", feed.FeedName, feed.Url, feed.UserName)
	}

	return nil
}
