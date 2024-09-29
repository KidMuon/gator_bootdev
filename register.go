package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kidmuon/gator_bootdev/internal/database"
	"time"
)

func handlerRegister(state *State, command Command) error {
	if len(command.args) == 0 {
		return fmt.Errorf("register requires a username")
	}

	username := command.args[0]

	userToCreate := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := state.db.CreateUser(context.Background(), userToCreate)
	if err != nil {
		return fmt.Errorf("user already registered: %v")
	}

	state.cfg.SetUser(username)

	fmt.Printf("user %s created and registered\n", user.Name)
	fmt.Println(user)

	return nil
}
