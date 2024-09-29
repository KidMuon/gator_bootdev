package main

import (
	"fmt"
)

type Command struct {
	name string
	args []string
}

type handlerFunc func(*State, Command) error

type Commands struct {
	validCommands map[string]func(*State, Command) error
}

func (c *Commands) register(name string, f func(state *State, command Command) error) {
	c.validCommands[name] = f
}

func (c *Commands) run(state *State, cmd Command) error {
	cmdFunc, ok := c.validCommands[cmd.name]
	if !ok {
		return fmt.Errorf("invalid/unregistered command passed")
	}

	return cmdFunc(state, cmd)
}
