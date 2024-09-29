package main

import (
	"fmt"
)

func handlerLogin(state *State, command Command) error {
	if len(command.args) < 1 {
		return fmt.Errorf("login requires a username")
	}
	username := command.args[0]

	err := state.SetUser(username)
	if err != nil {
		return fmt.Errorf("error logging in user %s: %v", username, err)
	}

	fmt.Printf("user %s logged in\n", username)
	return nil
}
