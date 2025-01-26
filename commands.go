package main

import (
	"errors"
)

type command struct {
	Name string
	Args []string
}

// map of command names to their handler functions
type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	callback, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("Command not found")
	}

	return callback(s, cmd)
}
