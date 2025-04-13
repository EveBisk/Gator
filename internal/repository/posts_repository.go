package repository

import (
	"context"
	"database/sql"
	db "gator/internal/database"
	"gator/internal/domain"

	"github.com/google/uuid"
)

type PostsRepository struct {
	q *db.Queries
}

func NewPostRepository(q *db.Queries) *PostsRepository {
	return &PostsRepository{q: q}
}

func (r *PostsRepository) CreatePost(ctx context.Context, post domain.Post) error {
	_, err := r.q.CreatePost(ctx, db.CreatePostParams{
		ID:          uuid.New(),
		FeedID:      post.FeedID,
		Url:         post.Url,
		Title:       post.Title,
		Description: toNullString(post.Description),
		PublishedAt: post.PublishedAt,
	})
	return err
}

func (r *PostsRepository) GetPostsForUser(ctx context.Context, userID uuid.UUID, limit int) ([]domain.Post, error) {
	posts, err := r.q.GetPostsForUser(ctx, db.GetPostsForUserParams{
		UserID: userID,
		Limit:  int32(limit),
	})
	if err != nil {
		return []domain.Post{}, err
	}
	return toDomainPosts(posts), nil
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}

func toDomainPosts(dbPosts []db.GetPostsForUserRow) []domain.Post {
	domain_list := []domain.Post{}

	for _, ff := range dbPosts {
		domain_list = append(domain_list, toDomainPost(ff))
	}
	return domain_list
}

func toDomainPost(dbPost db.GetPostsForUserRow) domain.Post {
	return domain.Post{
		ID:          dbPost.ID,
		Title:       dbPost.Title,
		Url:         dbPost.Url,
		Description: dbPost.Description.String,
		PublishedAt: dbPost.PublishedAt,
		FeedID:      dbPost.FeedID,
		FeedName:    dbPost.FeedName,
	}
}
