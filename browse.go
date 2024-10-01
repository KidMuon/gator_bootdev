package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kidmuon/gator_bootdev/internal/database"
)

func handlerBrowse(state *State, command Command, user database.User) error {
	var browseLimit int
	var err error
	if len(command.args) < 1 {
		browseLimit = 2
	} else {
		browseLimit, err = strconv.Atoi(command.args[0])
		if err != nil {
			return fmt.Errorf("problem converting passed limit to int: %v", err)
		}
	}

	postsToGet := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(browseLimit),
	}

	userPosts, err := state.db.GetPostsForUser(context.Background(), postsToGet)
	if err != nil {
		return fmt.Errorf("problem getting posts for user: %v", err)
	}

	for _, userPost := range userPosts {
		fmt.Printf("Title: %s\n", userPost.Title)
		fmt.Printf("Link: %s\n", userPost.Url.String)
		fmt.Printf("Description: %s\n", userPost.Description.String)
	}

	return nil
}
