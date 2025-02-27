package main

import (
	"fmt"
	"gator/internal/config"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error while reading config\n%w", err)
	}

	state_str := &state{
		cfg: cfg,
	}

	commands := commands{
		commandsMap: make(map[string]func(*state, command) error),
	}
	commands.populateCommandsMap()

	input := os.Args[1:]

	if len(input) < 2 {
		fmt.Printf("Expected a command name and its arguments")
		os.Exit(1)
	}

	cmd := command{
		name: input[0],
		args: input[1:],
	}

	err = commands.run(state_str, cmd)

	if err != nil {
		fmt.Println("Error while executing the command:", err)
		os.Exit(1)
	}
}
