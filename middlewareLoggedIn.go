package main

import (
	"context"
	"fmt"
	"github.com/kidmuon/gator_bootdev/internal/database"
)

func middlewareLoggedIn(handler func(state *State, command Command, user database.User) error) func(*State, Command) error {

	return func(state *State, command Command) error {
		user, err := state.db.GetUser(context.Background(), state.cfg.Username)
		if err != nil {
			return fmt.Errorf("error with user. not logged in or not in database: %v", err)
		}

		return handler(state, command, user)
	}
}
