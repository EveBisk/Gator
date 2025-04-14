package repository

import (
	"context"
	"time"

	db "github.com/EveBisk/gator/internal/database"
	"github.com/EveBisk/gator/internal/domain"

	"github.com/google/uuid"
)

type FeedFollowRepository struct {
	q *db.Queries
}

func NewFeedFollowRepository(q *db.Queries) *FeedFollowRepository {
	return &FeedFollowRepository{q: q}
}

func retrievedFeedToDomain(retrievedFeed db.GetFeedFollowsForUserRow) domain.FeedFollow {
	return domain.FeedFollow{
		ID:        retrievedFeed.ID,
		CreatedAt: retrievedFeed.CreatedAt,
		UpdatedAt: retrievedFeed.UpdatedAt,
		UserID:    retrievedFeed.UserID,
		FeedID:    retrievedFeed.FeedID,
		UserName:  retrievedFeed.UserName,
		FeedName:  retrievedFeed.FeedName,
	}
}

func retrievedFeedsToDomain(retrievedFeeds []db.GetFeedFollowsForUserRow) []domain.FeedFollow {
	domain_list := []domain.FeedFollow{}

	for _, ff := range retrievedFeeds {
		domain_list = append(domain_list, retrievedFeedToDomain(ff))
	}
	return domain_list
}

func toDomainFeedFollow(ff db.CreateFeedFollowRow) domain.FeedFollow {
	return domain.FeedFollow{
		ID:        ff.ID,
		FeedID:    ff.FeedID,
		UserID:    ff.UserID,
		CreatedAt: ff.CreatedAt,
		UpdatedAt: ff.UpdatedAt,
		UserName:  ff.UserName,
		FeedName:  ff.FeedName,
	}
}

func toDomainFeedFollows(ffs []db.CreateFeedFollowRow) []domain.FeedFollow {
	domain_list := []domain.FeedFollow{}

	for _, ff := range ffs {
		domain_list = append(domain_list, toDomainFeedFollow(ff))
	}
	return domain_list
}

func (r *FeedFollowRepository) CreateFeedFollow(ctx context.Context, feedId, userId uuid.UUID) (domain.FeedFollow, error) {
	feed_follow, err := r.q.CreateFeedFollow(ctx, db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userId,
		FeedID:    feedId,
	})
	if err != nil {
		return domain.FeedFollow{}, err
	}
	return toDomainFeedFollow(feed_follow), err
}

func (r *FeedFollowRepository) GetFeedFollowsForUser(ctx context.Context, userId uuid.UUID) ([]domain.FeedFollow, error) {
	feed_follows, err := r.q.GetFeedFollowsForUser(ctx, userId)
	if err != nil {
		return []domain.FeedFollow{}, err
	}
	return retrievedFeedsToDomain(feed_follows), err
}

func (r *FeedFollowRepository) RemoveFollowByUserURL(ctx context.Context, feedUrl string, userId uuid.UUID) error {
	return r.q.RemoveFollowByUserURL(ctx, db.RemoveFollowByUserURLParams{
		UserID: userId,
		Url:    feedUrl,
	})
}
