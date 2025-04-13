package main

import (
	"fmt"
	"gator/internal/config"
	"gator/internal/domain"
	"gator/internal/service"
	"log"
	"time"
)

func handlerAggFeed(s *config.State, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	time_between_reqs := cmd.args[0]
	time_duration, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("couldn't parse time duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s", time_between_reqs)

	ticker := time.NewTicker(time_duration)
	for ; ; <-ticker.C {
		service.ScrapeFeeds(s)
	}
}

func handlerAddFeed(s *config.State, cmd command, user domain.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	name := cmd.args[0]
	feedUrl := cmd.args[1]

	feed, err := s.Repos.FeedRepo.CreateFeed(s.Ctx, feedUrl, name, user.ID)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	_, err = s.Repos.FeedFollowRepo.CreateFeedFollow(s.Ctx, feed.ID, feed.UserID)
	if err != nil {
		log.Printf("Error while adding feed follow")
	}

	fmt.Print("Feed entry created:\n")
	fmt.Printf("%+v", feed)
	return nil
}

func handlerGetAllFeeds(s *config.State, cmd command) error {
	feeds, err := service.GetFeedsWithUsers(s)
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		feed.PrintFeed()
	}
	return nil
}
