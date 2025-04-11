package main

import (
	"gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.dbQueries.GetUserByName(s.ctx, s.cfg.Current_user_name)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
