package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
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

func printFeed(feed RSSFeed) {
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
