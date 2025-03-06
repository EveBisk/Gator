package main

import (
	"context"
	"database/sql"
	"errors"
	"gator/internal/config"
	"gator/internal/database"
)

type state struct {
	cfg       *config.Config
	db        *sql.DB
	dbQueries *database.Queries
	ctx       context.Context
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
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerUsers)
	c.register("agg", handlerFetchFeed)
	c.register("addfeed", handlerAddFeed)
	c.register("feeds", handlerGetAllFeeds)
}

func getNoArgsCommands() []string {
	return []string{"reset", "users", "agg", "feeds"}
}
