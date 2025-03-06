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

	fmt.Print("Feed entry created:\n")
	fmt.Printf("%+v", feed)
	return nil
}

func handlerGetAllFeeds(s *state, cmd command) error {
	feeds, err := s.dbQueries.GetAllFeeds(s.ctx)
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		usr, err := s.dbQueries.GetUserById(s.ctx, feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get feed user: %w", err)
		}
		printFeed(dbFeedToFeed(feed, usr.Name))
	}
	return nil
}
