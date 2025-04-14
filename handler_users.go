package main

import (
	"fmt"
	"gator/internal/config"
)

func handlerGetUsers(s *config.State, cmd command) error {
	usrs, err := s.Repos.UserRepo.GetUsers(s.Ctx)
	if err != nil {
		return fmt.Errorf("couldn't get users: %w", err)
	}

	current_user, err := s.Cfg.GetCurrentUser()

	if err != nil {
		return err
	}

	fmt.Print("Users currently in db:\n")
	for _, u := range usrs {
		if u.Name == current_user {
			fmt.Printf("* %s (current)\n", u.Name)
			continue
		}
		fmt.Printf("* %s\n", u.Name)
	}
	return nil
}
