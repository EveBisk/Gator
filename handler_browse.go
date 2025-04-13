package main

import (
	"fmt"
	"gator/internal/config"
	"gator/internal/domain"
	"strconv"
)

func handlerBrowse(s *config.State, cmd command, user domain.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.args[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.Repos.PostsRepo.GetPostsForUser(s.Ctx, user.ID, limit)
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
