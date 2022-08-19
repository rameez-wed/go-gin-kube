package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	mockdb "github.com/go-gin-kube/db/mock"
	db "github.com/go-gin-kube/db/sqlc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAuthorAPI(t *testing.T) {
	author := randomAuthor()

	testCases := []struct {
		name          string
		authorID      int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:     "OK",
			authorID: author.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAuthorById(gomock.Any(), gomock.Eq(author.ID)).
					Times(1).
					Return(author, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatcherAuthor(t, recorder.Body, author)
			},
		},
		// TODO: add more cases
		{
			name:     "NotFound",
			authorID: author.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAuthorById(gomock.Any(), gomock.Eq(author.ID)).
					Times(1).
					Return(db.Author{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:     "InternalServerError",
			authorID: author.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAuthorById(gomock.Any(), gomock.Eq(author.ID)).
					Times(1).
					Return(db.Author{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:     "InvalidInput",
			authorID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAuthorById(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			// build stubs

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/authors/%d", tc.authorID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

		// check response
	}

}

func randomAuthor() db.Author {
	return db.Author{
		ID:   1,
		Name: faker.Name(),
	}
}

func requireBodyMatcherAuthor(t *testing.T, body *bytes.Buffer, author db.Author) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	var returnedAuthor db.Author
	err = json.Unmarshal(data, &returnedAuthor)
	require.NoError(t, err)
	require.Equal(t, author, returnedAuthor)
}
