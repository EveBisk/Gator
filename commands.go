package main

import (
	"context"
	"database/sql"
	"errors"
	"gator/internal/config"
	"gator/internal/database"
	repo "gator/internal/repository"
)

type state struct {
	cfg       *config.Config
	db        *sql.DB
	dbQueries *database.Queries
	ctx       context.Context
	repos     repositories
}

type repositories struct {
	userRepo       *repo.UserRepository
	feedRepo       *repo.FeedRepository
	feedFollowRepo *repo.FeedFollowRepository
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
	c.register("users", handlerGetUsers)
	c.register("agg", handlerFetchFeed)
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("feeds", handlerGetAllFeeds)
	c.register("follow", middlewareLoggedIn(handlerFollow))
	c.register("following", middlewareLoggedIn(handlerGetFollowing))
	c.register("unfollow", middlewareLoggedIn(handlerDeleteFollow))
}

func getNoArgsCommands() []string {
	return []string{"reset", "users", "feeds", "following"}
}
