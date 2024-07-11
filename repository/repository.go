package repository

import (
	"context"

	"github.com/YugenDev/go-platzi-advanced-two/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertPost(ctx context.Context, post *models.Posts) error
	GetPostById(ctx context.Context, id string) (*models.Posts, error)
	UpdatePost(ctx context.Context, post *models.Posts) error
	DeletePost(ctx context.Context, id string, userId string) error
	ListPosts(ctx context.Context, page uint64) ([]*models.Posts, error)
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func Close() error {
	return implementation.Close()
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func InsertPost(ctx context.Context, post *models.Posts) error {
	return implementation.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id string) (*models.Posts, error) {
	return implementation.GetPostById(ctx, id)
}

func UpdatePost(ctx context.Context, post *models.Posts) error {
	return implementation.UpdatePost(ctx, post)
}
func DeletePost(ctx context.Context, id string, userId string) error {
	return implementation.DeletePost(ctx, id, userId)
}

func ListPosts(ctx context.Context, page uint64) ([]*models.Posts, error) {
	return implementation.ListPosts(ctx, page)
}
