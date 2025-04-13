package service

import (
	"context"
	"fmt"
	"gator/internal/domain"
	"gator/internal/repository"

	"github.com/google/uuid"
)

func CreateFeedFollowForUser(
	ctx context.Context,
	feedRepo *repository.FeedRepository,
	feedFollowRepo *repository.FeedFollowRepository,
	feedUrl string,
	userId uuid.UUID,
) (domain.FeedFollow, error) {
	feed, err := feedRepo.GetFeedIdFromURL(ctx, feedUrl)
	if err != nil {
		return domain.FeedFollow{}, err
	}

	feed_follow, err := feedFollowRepo.CreateFeedFollow(ctx, feed.ID, userId)
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
