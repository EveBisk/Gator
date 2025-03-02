package main

import (
	"database/sql"
	"gator/internal/config"
	"gator/internal/database"
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

	state_str := &state{
		cfg:       cfg,
		db:        db,
		dbQueries: database.New(db),
	}

	commands := commands{
		commandsMap: make(map[string]func(*state, command) error),
	}
	commands.populateCommandsMap()

	input := os.Args[1:]

	switch {
	case len(input) == 0:
		log.Fatal("Usage: cli <command> [args...]")
	case len(input) == 1 && !slices.Contains(getNoArgsCommands(), input[0]):
		log.Fatal("Usage: cli <command> [args...]")
	}

	err = commands.run(state_str, command{name: input[0], args: input[1:]})

	if err != nil {
		log.Fatal(err)
	}
}
