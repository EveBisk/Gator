package domain

import (
	"fmt"
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
	UserID    uuid.UUID
}

func (feed RSSFeed) PrintRSSFeed() {
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

func (feed Feed) PrintFeed() {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserName:        %s\n", feed.UserName)
}
