package main

import (
	"gator/internal/config"
	"gator/internal/domain"
)

func middlewareLoggedIn(handler func(s *config.State, cmd command, user domain.User) error) func(*config.State, command) error {
	return func(s *config.State, cmd command) error {
		user, err := s.Repos.UserRepo.GetUserByName(s.Ctx, s.Cfg.Current_user_name)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
