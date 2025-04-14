package main

import (
	"fmt"
	"gator/internal/config"
)

func handlerReset(s *config.State, cmd command) error {
	err := s.Repos.UserRepo.DeleteUsers(s.Ctx)
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}
