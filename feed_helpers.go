package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"gator/internal/database"
	"html"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type Feed struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserName  string
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

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

	feed := RSSFeed{}
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

func printRSSFeed(feed RSSFeed) {
	channel := feed.Channel
	fmt.Printf(" * Title:      %v\n", channel.Title)
	fmt.Printf(" * Link:    %v\n", channel.Link)
	fmt.Printf(" * Description:    %v\n", channel.Description)
	fmt.Print(" * Items:")
	for _, item := range channel.Item {
		fmt.Printf("\t * Title:      %v\n", item.Title)
		fmt.Printf("\t * Link:    %v\n", item.Link)
		fmt.Printf("\t * Description:    %v\n", item.Description)
	}
}

func printFeed(feed Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserName:        %s\n", feed.UserName)
}

func dbFeedToFeed(dbFeed database.Feed, userName string) Feed {
	return Feed{
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		ID:        dbFeed.ID,
		UserName:  userName,
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.dbQueries.GetNextFeedToFetch(s.ctx)
	if err != nil {
		return fmt.Errorf("error retrieving next feed url %w", err)
	}

	fmt.Printf("Sending request to %s", feed.Url)

	err = s.dbQueries.MarkFeedFetched(s.ctx, feed.ID)
	if err != nil {
		log.Printf("error marking feed %v as fetched", feed.ID)
	}

	feed_content, err := fetchFeed(s.ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed content %w", err)
	}

	printRSSFeed(*feed_content)
	return nil
}
