package main

import (
	"errors"
	"gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandsMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandsMap[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	fun, ok := c.commandsMap[cmd.name]

	if !ok {
		return errors.New("command does not exist")
	}

	return fun(s, cmd)
}

func (c *commands) populateCommandsMap() {
	c.register("login", handlerLogin)
}
