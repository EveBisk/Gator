package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID
	Title       string
	Url         string
	Description string
	PublishedAt time.Time
	FeedID      uuid.UUID
	FeedName    string
}
