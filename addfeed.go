package main

import (
	"context"
	"fmt"

	"github.com/kidmuon/gator_bootdev/internal/database"
)

func handlerAddFeed(state *State, command Command) error {
	if len(command.args) < 2 {
		return fmt.Errorf("addfeed requires two arguments, name and url")
	}

	feedName := command.args[0]
	feedUrl := command.args[1]
	user, err := state.db.GetUser(context.Background(), state.cfg.Username)
	if err != nil {
		return fmt.Errorf("unable to get current user's information from database: %v", err)
	}

	feedToCreate := database.CreateFeedParams{
		Name:   feedName,
		Url:    feedUrl,
		UserID: user.ID,
	}

	feed, err := state.db.CreateFeed(context.Background(), feedToCreate)
	if err != nil {
		return fmt.Errorf("unable to create feed: %v", err)
	}

	fmt.Println(feed)

	return nil
}
