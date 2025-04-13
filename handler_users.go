package main

import (
	"fmt"
)

func handlerGetUsers(s *state, cmd command) error {
	usrs, err := s.repos.userRepo.GetUsers(s.ctx)
	if err != nil {
		return fmt.Errorf("couldn't get users: %w", err)
	}

	current_user, err := s.cfg.GetCurrentUser()

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
