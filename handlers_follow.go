package main

import (
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	feedUrl := cmd.args[0]

	feed, err := s.dbQueries.GetFeedIdFromURL(s.ctx, feedUrl)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	feed_follow, err := s.dbQueries.CreateFeedFollow(s.ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed - follow: %w", err)
	}

	fmt.Print("Feed - follow entry created:\n")
	fmt.Printf("%+v", feed_follow)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feed_follows, err := s.dbQueries.GetFeedFollowsForUser(s.ctx, user.ID)
	if err != nil {
		return fmt.Errorf("couldn't fetch feeds for user: %w", err)
	}

	fmt.Printf("User %s is following the feeds:\n", user.Name)
	for _, feed := range feed_follows {
		fmt.Printf("\t* %s\n", feed.FeedName)
	}
	return nil
}

func handlerDeleteFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	feedUrl := cmd.args[0]

	err := s.dbQueries.RemoveFollowByUserURL(s.ctx, database.RemoveFollowByUserURLParams{
		UserID: user.ID,
		Url:    feedUrl,
	})
	if err != nil {
		return fmt.Errorf("couldn't fetch feeds for user: %w", err)
	}

	fmt.Print("Feed successfully removed from your follow list")

	return nil
}
