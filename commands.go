package main

import (
	"errors"
	"gator/internal/config"
)

type command struct {
	name string
	args []string
}

type commands struct {
	commandsMap map[string]func(*config.State, command) error
}

func (c *commands) register(name string, f func(*config.State, command) error) {
	c.commandsMap[name] = f
}

func (c *commands) run(s *config.State, cmd command) error {
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
	c.register("users", handlerGetUsers)
	c.register("agg", handlerAggFeed)
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("feeds", handlerGetAllFeeds)
	c.register("follow", middlewareLoggedIn(handlerFollow))
	c.register("following", middlewareLoggedIn(handlerGetFollowing))
	c.register("unfollow", middlewareLoggedIn(handlerDeleteFollow))
	c.register("browse", middlewareLoggedIn(handlerBrowse))
}

func getNoArgsCommands() []string {
	return []string{"reset", "users", "feeds", "following"}
}
