package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	testDb  *MemDB
	handler BookHandler
)

var (
	testBookID   = "a123"
	testAuthorID = "b123"
)

func TestMain(m *testing.M) {
	testDb = NewMemDB(testAuthorID, testBookID)
	addSampleAuthor()
	handler = NewBookHandler(testDb)

	m.Run()
}

func addSampleAuthor() {
	now := time.Now()
	testDb.authors["aass"] = &Author{
		ID:        "aass",
		Firstname: "Alan",
		Lastname:  "Donovan",
		CreatedOn: now,
		UpdatedOn: now,
		DeletedOn: nil,
	}
	testDb.authors["aabb"] = &Author{
		ID:        "aabb",
		Firstname: "Brian",
		Lastname:  "Kernighan",
		CreatedOn: now,
		UpdatedOn: now,
		DeletedOn: nil,
	}
}

func TestPostBook(t *testing.T) {
	body := `{
        "title": "The Go Programming Language",
        "authorId": "aass",
        "totalPage": 380
    }`
	req, _ := http.NewRequest(http.MethodPost, "/book", strings.NewReader(body))
	res := httptest.NewRecorder()

	handler.handlePostBook(res, req)

	resId := make(map[string]string)

	json.Unmarshal(res.Body.Bytes(), &resId)
	require.Equal(t, http.StatusOK, res.Code)
	require.Equal(t, testBookID, resId["bookId"])
}

func TestGetAllBooks(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/book", nil)
	res := httptest.NewRecorder()

	handler.handleGetBooks(res, req)

	books := make(map[string][]*Book)
	json.Unmarshal(res.Body.Bytes(), &books)

	require.Equal(t, 1, len(books), "length of books return")
	require.Equal(t, http.StatusOK, res.Code, "response code")
	require.Equal(t, testBookID, books["books"][0].ID, "mock first item's id")
}
