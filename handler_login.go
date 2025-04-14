package main

import (
	"fmt"

	"github.com/EveBisk/gator/internal/config"
)

func handlerLogin(s *config.State, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	username := cmd.args[0]
	usr, err := s.Repos.UserRepo.GetUserByName(s.Ctx, username)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.Cfg.SetUser(usr.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}
