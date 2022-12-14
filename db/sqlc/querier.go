// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
)

type Querier interface {
	CreateAuthor(ctx context.Context, name string) (Author, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	DeleteAuthor(ctx context.Context, id int64) error
	DeletePost(ctx context.Context, id int64) error
	GetAllPostsForAuthor(ctx context.Context, arg GetAllPostsForAuthorParams) ([]Post, error)
	GetAuthorById(ctx context.Context, id int64) (Author, error)
	GetPostById(ctx context.Context, id int64) (Post, error)
	ListAuthors(ctx context.Context, arg ListAuthorsParams) ([]Author, error)
	ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error)
	UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
}

var _ Querier = (*Queries)(nil)
