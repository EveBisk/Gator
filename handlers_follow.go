package main

import (
	"fmt"
	"gator/internal/domain"
	"gator/internal/service"
)

func handlerFollow(s *state, cmd command, user domain.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	feedUrl := cmd.args[0]
	feed_follow, err := service.CreateFeedFollowForUser(
		s.ctx,
		s.repos.feedRepo,
		s.repos.feedFollowRepo,
		feedUrl,
		user.ID,
	)

	if err != nil {
		return fmt.Errorf("couldn't create feed - follow: %w", err)
	}

	fmt.Print("Feed - follow entry created:\n")
	fmt.Printf("%+v", feed_follow)
	return nil
}

func handlerGetFollowing(s *state, cmd command, user domain.User) error {
	feed_follows, err := s.repos.feedFollowRepo.GetFeedFollowsForUser(s.ctx, user.ID)
	if err != nil {
		return fmt.Errorf("couldn't fetch feeds for user: %w", err)
	}

	fmt.Printf("User %s is following the feeds:\n", user.Name)
	service.PrintFeedNames(feed_follows)
	return nil
}

func handlerDeleteFollow(s *state, cmd command, user domain.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	feedUrl := cmd.args[0]
	err := s.repos.feedFollowRepo.RemoveFollowByUserURL(
		s.ctx,
		feedUrl,
		user.ID,
	)

	if err != nil {
		return fmt.Errorf("couldn't fetch feeds for user: %w", err)
	}

	fmt.Print("Feed successfully removed from your follow list")
	return nil
}
