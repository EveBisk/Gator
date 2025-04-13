package service

import (
	"gator/internal/config"
	"gator/internal/domain"
	"time"

	"github.com/google/uuid"
)

// Should do some error handling here
func SavePostsofRSSFeed(s *config.State, rssFeed domain.RSSFeed, feedId uuid.UUID) {
	channel := rssFeed.Channel

	for _, post := range channel.Item {
		pubdate, _ := time.Parse(time.RFC3339, post.PubDate)
		s.Repos.PostsRepo.CreatePost(s.Ctx, domain.Post{
			Title:       post.Title,
			Url:         post.Link,
			PublishedAt: pubdate,
			Description: post.Description,
			FeedID:      feedId,
		})
	}
}
