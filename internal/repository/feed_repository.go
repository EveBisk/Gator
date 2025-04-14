package repository

import (
	"context"
	"time"

	"github.com/EveBisk/gator/internal/database"

	"github.com/EveBisk/gator/internal/domain"

	"github.com/google/uuid"
)

type FeedRepository struct {
	q *database.Queries
}

func NewFeedRepository(q *database.Queries) *FeedRepository {
	return &FeedRepository{q: q}
}

func dbFeedToFeed(dbFeed database.Feed) domain.Feed {
	return domain.Feed{
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		ID:        dbFeed.ID,
		UserID:    dbFeed.UserID,
		UserName:  "",
	}
}

func dbFeedsToFeeds(dbFeeds []database.Feed) []domain.Feed {
	feeds := []domain.Feed{}
	for _, feed := range dbFeeds {
		feeds = append(feeds, dbFeedToFeed(feed))
	}
	return feeds
}

func (r *FeedRepository) GetNextFeedToFetch(ctx context.Context) (domain.Feed, error) {
	feed, err := r.q.GetNextFeedToFetch(ctx)
	if err != nil {
		return domain.Feed{}, err
	}
	return dbFeedToFeed(feed), nil
}

func (r *FeedRepository) MarkFeedFetched(ctx context.Context, id uuid.UUID) error {
	return r.q.MarkFeedFetched(ctx, id)
}

func (r *FeedRepository) GetAllFeeds(ctx context.Context) ([]domain.Feed, error) {
	feeds, err := r.q.GetAllFeeds(ctx)
	if err != nil {
		return []domain.Feed{}, err
	}
	return dbFeedsToFeeds(feeds), nil
}

func (r *FeedRepository) CreateFeed(ctx context.Context, url, name string, userId uuid.UUID) (domain.Feed, error) {
	feed, err := r.q.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       url,
		Name:      name,
		UserID:    userId,
	})
	if err != nil {
		return domain.Feed{}, err
	}
	return dbFeedToFeed(feed), nil
}

func (r *FeedRepository) GetFeedIdFromURL(ctx context.Context, url string) (domain.Feed, error) {
	feed, err := r.q.GetFeedIdFromURL(ctx, url)
	if err != nil {
		return domain.Feed{}, err
	}
	return dbFeedToFeed(feed), nil
}
