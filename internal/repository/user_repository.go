package repository

import (
	"context"
	db "gator/internal/database"
	"gator/internal/domain"
	"time"

	"github.com/google/uuid"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func convertToDomainUser(dbUser db.User) domain.User {
	return domain.User{
		ID:   dbUser.ID,
		Name: dbUser.Name,
	}
}

func convertToDomainUsers(dbUsers []db.User) []domain.User {
	domain_users := []domain.User{}

	for _, user := range dbUsers {
		domain_users = append(domain_users, convertToDomainUser(user))
	}
	return domain_users
}

func (r *UserRepository) GetUserByName(ctx context.Context, name string) (domain.User, error) {
	user, err := r.q.GetUserByName(ctx, name)
	if err != nil {
		return domain.User{}, err
	}
	return convertToDomainUser(user), nil
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	users, err := r.q.GetUsers(ctx)
	if err != nil {
		return []domain.User{}, err
	}
	return convertToDomainUsers(users), nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	user, err := r.q.GetUserById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return convertToDomainUser(user), nil
}

func (r *UserRepository) CreateUser(ctx context.Context, name string) (domain.User, error) {
	new_user_params := db.CreateUserParams{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	usr, err := r.q.CreateUser(ctx, new_user_params)
	if err != nil {
		return domain.User{}, err
	}

	return convertToDomainUser(usr), nil
}

func (r *UserRepository) DeleteUsers(ctx context.Context) error {
	return r.q.DeleteUsers(ctx)
}
