package service

import (
	"fmt"

	"github.com/EveBisk/gator/internal/config"

	"github.com/EveBisk/gator/internal/domain"

	"github.com/google/uuid"
)

func CreateFeedFollowForUser(
	s *config.State,
	feedUrl string,
	userId uuid.UUID,
) (domain.FeedFollow, error) {
	feed, err := s.Repos.FeedRepo.GetFeedIdFromURL(s.Ctx, feedUrl)
	if err != nil {
		return domain.FeedFollow{}, err
	}

	feed_follow, err := s.Repos.FeedFollowRepo.CreateFeedFollow(s.Ctx, feed.ID, userId)
	if err != nil {
		return domain.FeedFollow{}, err
	}
	return feed_follow, err
}

func PrintFeedNames(feed_follows []domain.FeedFollow) {
	for _, feed := range feed_follows {
		fmt.Printf("\t* %s\n", feed.FeedName)
	}
}
