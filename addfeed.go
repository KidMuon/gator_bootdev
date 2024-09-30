package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kidmuon/gator_bootdev/internal/database"
)

func handlerAddFeed(state *State, command Command, user database.User) error {
	if len(command.args) < 2 {
		return fmt.Errorf("addfeed requires two arguments, name and url")
	}

	feedName := command.args[0]
	feedUrl := command.args[1]

	feedToCreate := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	}

	feed, err := state.db.CreateFeed(context.Background(), feedToCreate)
	if err != nil {
		return fmt.Errorf("unable to create feed: %v", err)
	}

	err = handlerFollow(state, Command{name: "follow", args: []string{feedUrl}}, user)
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
