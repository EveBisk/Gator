package main

import (
	"context"
	"database/sql"
	"gator/internal/config"
	"gator/internal/database"
	"gator/internal/repository"
	"log"
	"os"
	"slices"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error while reading config\n%v", err)
	}

	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}

	dbQueries := database.New(db)
	// I am using config.State as a form of depedency injection in the different parts of the code
	state := &config.State{
		Cfg:       cfg,
		Db:        db,
		DbQueries: dbQueries,
		Ctx:       context.Background(),
		Repos: config.Repositories{
			UserRepo:       repository.NewUserRepository(dbQueries),
			FeedRepo:       repository.NewFeedRepository(dbQueries),
			FeedFollowRepo: repository.NewFeedFollowRepository(dbQueries),
			PostsRepo:      repository.NewPostRepository(dbQueries),
		},
	}

	commands := commands{
		commandsMap: make(map[string]func(*config.State, command) error),
	}
	commands.populateCommandsMap()

	input := os.Args[1:]

	switch {
	case len(input) == 0:
		log.Fatal("Usage: cli <command> [args...]")
	case len(input) == 1 && !slices.Contains(getNoArgsCommands(), input[0]):
		log.Fatal("Usage: cli <command> [args...]")
	}

	err = commands.run(state, command{name: input[0], args: input[1:]})

	if err != nil {
		log.Fatal(err)
	}
}
