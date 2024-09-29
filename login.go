package main

import (
	"context"
	"fmt"
)

func handlerLogin(state *State, command Command) error {
	if len(command.args) < 1 {
		return fmt.Errorf("login requires a username")
	}
	username := command.args[0]

	_, err := state.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("username %s not registered with database", username)
	}

	err = state.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error logging in user %s: %v", username, err)
	}

	fmt.Printf("user %s logged in\n", username)
	return nil
}
