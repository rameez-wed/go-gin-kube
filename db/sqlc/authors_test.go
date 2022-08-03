package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func createRandomAuthor(t *testing.T) Author {
	randomAuthor := Author{Name: faker.FirstName()}
	author, err := testQueries.CreateAuthor(context.Background(), randomAuthor.Name)
	require.NoError(t, err)
	require.NotEmpty(t, author)
	require.NotEmpty(t, author.ID)
	require.Equal(t, randomAuthor.Name, author.Name)
	return author
}

func TestCreateAuthor(t *testing.T) {
	createRandomAuthor(t)
}

func TestGetAuthor(t *testing.T) {
	randomAuthor := createRandomAuthor(t)
	require.NotEmpty(t, randomAuthor)
	author, err := testQueries.GetAuthorById(context.Background(), randomAuthor.ID)
	require.NoError(t, err)
	require.NotEmpty(t, author)
	require.NotEmpty(t, author.ID)
	require.Equal(t, randomAuthor.ID, author.ID)
	require.Equal(t, randomAuthor.Name, author.Name)
}

func TestListAuthor(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAuthor(t)
	}
	listAuthorArgs := ListAuthorsParams{Limit: 5, Offset: 5}
	authors, err := testQueries.ListAuthors(context.Background(), listAuthorArgs)
	require.NoError(t, err)
	require.Len(t, authors, 5)
	for _, author := range authors {
		require.NotEmpty(t, author)
	}
}

func TestUpdateAuthor(t *testing.T) {
	randomAuthor := createRandomAuthor(t)
	require.NotEmpty(t, randomAuthor)
	updatedName := "John"
	updateAuthorParams := UpdateAuthorParams{randomAuthor.ID, updatedName}
	author, err := testQueries.UpdateAuthor(context.Background(), updateAuthorParams)
	require.NoError(t, err)
	require.NotEmpty(t, author)
	require.Equal(t, author.Name, updateAuthorParams.Name)
}

func TestDeleteAuthor(t *testing.T) {
	randomAuthor := createRandomAuthor(t)
	require.NotEmpty(t, randomAuthor)
	err := testQueries.DeleteAuthor(context.Background(), randomAuthor.ID)
	require.NoError(t, err)
	author, err := testQueries.GetAuthorById(context.Background(), randomAuthor.ID)
	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows)
	require.Empty(t, author)
}
