package service

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/EveBisk/gator/internal/config"

	"github.com/EveBisk/gator/internal/domain"
)

func ScrapeFeeds(s *config.State) error {
	// Depedencies
	feedRepo := s.Repos.FeedRepo

	feed, err := feedRepo.GetNextFeedToFetch(s.Ctx)
	if err != nil {
		return fmt.Errorf("error retrieving next feed url %w", err)
	}

	fmt.Printf("Sending request to %s", feed.Url)

	err = feedRepo.MarkFeedFetched(s.Ctx, feed.ID)
	if err != nil {
		return err
	}

	feed_content, err := fetchFeed(s.Ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed content %w", err)
	}

	// We are going to save the feeds later
	SavePostsofRSSFeed(s, feed_content, feed.ID)
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (domain.RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return domain.RSSFeed{}, fmt.Errorf("error creating new rss feed request: %v", err)
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
		return domain.RSSFeed{}, fmt.Errorf("request failed: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.RSSFeed{}, fmt.Errorf("error reading response body: %v", err)
	}

	feed := domain.RSSFeed{}
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return domain.RSSFeed{}, fmt.Errorf("error unmarshalling response body: %v", err)
	}

	for _, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	return feed, nil
}

func GetFeedsWithUsers(s *config.State) ([]domain.Feed, error) {
	feeds, err := s.Repos.FeedRepo.GetAllFeeds(s.Ctx)
	if err != nil {
		return []domain.Feed{}, fmt.Errorf("couldn't get feeds: %w", err)
	}

	for i := range feeds {
		feed := &feeds[i]
		usr, err := s.Repos.UserRepo.GetUserByID(s.Ctx, feed.UserID)
		if err != nil {
			return []domain.Feed{}, fmt.Errorf("couldn't get feed user: %w", err)
		}
		feed.UserName = usr.Name
	}
	return feeds, nil
}
