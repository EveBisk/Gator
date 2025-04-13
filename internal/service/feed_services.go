package service

import (
	"context"
	"encoding/xml"
	"fmt"
	"gator/internal/domain"
	"gator/internal/repository"
	"html"
	"io"
	"net/http"
	"time"
)

func ScrapeFeeds(feedRepo *repository.FeedRepository, ctx context.Context) error {
	feed, err := feedRepo.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving next feed url %w", err)
	}

	fmt.Printf("Sending request to %s", feed.Url)

	err = feedRepo.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return err
	}

	feed_content, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed content %w", err)
	}

	// We are going to save the feeds later
	feed_content.PrintRSSFeed()
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*domain.RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating new rss feed request: %v", err)
	}
	req.Header.Set("User-Agent", "gator")

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	feed := domain.RSSFeed{}
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v", err)
	}

	for _, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	return &feed, nil
}

func GetFeedsWithUsers(feedRepo *repository.FeedRepository, userRepo *repository.UserRepository, ctx context.Context) ([]domain.Feed, error) {
	feeds, err := feedRepo.GetAllFeeds(ctx)
	if err != nil {
		return []domain.Feed{}, fmt.Errorf("couldn't get feeds: %w", err)
	}

	for i := range feeds {
		feed := &feeds[i]
		usr, err := userRepo.GetUserByID(ctx, feed.UserID)
		if err != nil {
			return []domain.Feed{}, fmt.Errorf("couldn't get feed user: %w", err)
		}
		feed.UserName = usr.Name
	}
	return feeds, nil
}
