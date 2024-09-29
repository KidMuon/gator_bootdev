package main

import (
	"context"
	"fmt"
)

func handlerReset(state *State, command Command) error {
	err := state.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("reset failed")
	}
	fmt.Println("reset successful")
	return nil
}
