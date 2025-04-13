package main

import (
	"fmt"
	"gator/internal/config"
)

func handlerRegister(s *config.State, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	usr, err := s.Repos.UserRepo.CreateUser(s.Ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.Cfg.SetUser(usr.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User has been created successfully:")
	usr.Print()
	return nil
}
