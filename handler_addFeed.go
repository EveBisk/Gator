package main

import (
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	name := cmd.args[0]
	feedUrl := cmd.args[1]

	user, err := s.dbQueries.GetUserByName(s.ctx, s.cfg.Current_user_name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	feed, err := s.dbQueries.CreateFeed(s.ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       feedUrl,
		Name:      name,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Printf("%+v", feed)
	return nil
}
