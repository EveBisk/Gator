package config

import (
	"context"
	"database/sql"
	"gator/internal/database"
	repo "gator/internal/repository"
)

type State struct {
	Cfg       *Config
	Db        *sql.DB
	DbQueries *database.Queries
	Ctx       context.Context
	Repos     Repositories
}

type Repositories struct {
	UserRepo       *repo.UserRepository
	FeedRepo       *repo.FeedRepository
	FeedFollowRepo *repo.FeedFollowRepository
	PostsRepo      *repo.PostsRepository
}
