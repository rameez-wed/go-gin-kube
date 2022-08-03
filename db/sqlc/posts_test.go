package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T, authorId int64) Post {
	randomPost := Post{Title: faker.Word(), Description: faker.Paragraph()}
	postParams := CreatePostParams{Title: randomPost.Title, Description: randomPost.Description, AuthorID: authorId}
	post, err := testQueries.CreatePost(context.Background(), postParams)
	require.NoError(t, err)
	require.NotEmpty(t, post)
	return post
}

func TestCreatePost(t *testing.T) {
	author := createRandomAuthor(t)
	createRandomPost(t, author.ID)
}

func TestGetPost(t *testing.T) {
	author := createRandomAuthor(t)
	post := createRandomPost(t, author.ID)
	fetchedPost, err := testQueries.GetPostById(context.Background(), post.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedPost)
	require.Equal(t, post.Title, fetchedPost.Title)
	require.Equal(t, post.Description, fetchedPost.Description)
	require.Equal(t, post.AuthorID, fetchedPost.AuthorID)
}

func TestDeletePost(t *testing.T) {
	author := createRandomAuthor(t)
	post := createRandomPost(t, author.ID)
	err := testQueries.DeletePost(context.Background(), post.ID)
	require.NoError(t, err)
	fetchedPost, err := testQueries.GetPostById(context.Background(), post.ID)
	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows)
	require.Empty(t, fetchedPost)
}

func TestListPosts(t *testing.T) {
	author := createRandomAuthor(t)
	for i := 0; i < 10; i++ {
		createRandomPost(t, author.ID)
	}
	listPostParams := ListPostsParams{Limit: 5, Offset: 5}
	posts, err := testQueries.ListPosts(context.Background(), listPostParams)
	require.NoError(t, err)
	require.NotEmpty(t, posts)
	require.Len(t, posts, 5)
	for _, post := range posts {
		require.NotEmpty(t, post)
	}
}

func TestUpdatePost(t *testing.T) {
	author := createRandomAuthor(t)
	randomPost := createRandomPost(t, author.ID)
	updatePostParams := UpdatePostParams{
		ID:          randomPost.ID,
		Title:       "updated title",
		Description: "updated description",
	}
	post, err := testQueries.UpdatePost(context.Background(), updatePostParams)
	require.NoError(t, err)
	require.NotEmpty(t, post)
	require.Equal(t, post.Title, updatePostParams.Title)
	require.Equal(t, post.Description, updatePostParams.Description)
}

func TestGetAllPostsForAuthor(t *testing.T) {
	author := createRandomAuthor(t)
	posts := make([]Post, 5)
	for i := 0; i < 5; i++ {
		posts = append(posts, createRandomPost(t, author.ID))
	}
	authorsPostsParams := GetAllPostsForAuthorParams{author.ID, 10, 0}
	authorsPosts, err := testQueries.GetAllPostsForAuthor(context.Background(), authorsPostsParams)
	require.NoError(t, err)
	require.NotEmpty(t, authorsPosts)
	require.Len(t, authorsPosts, 5)
}
