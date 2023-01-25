package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type BookDB interface {
	GetAll(ctx context.Context) ([]*Book, error)
	Add(ctx context.Context, book *Book) (string, error)

	AddAuthor(ctx context.Context, author *Author) (string, error)

	GetByID(ctx context.Context, id string) (*Book, *Author, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, field BookProp, value any) (*Book, error)
}

type BookHandler struct {
	db BookDB
}

func NewBookHandler(db BookDB) BookHandler {
	return BookHandler{db: db}
}

func (bh BookHandler) handleGetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := bh.db.GetAll(r.Context())
	if err != nil {
		JsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Json(w, BooksResponse{Books: books}, http.StatusOK)
}

type BooksResponse struct {
	Books []*Book `json:"books"`
}

func (bh BookHandler) handleGetBookByID(w http.ResponseWriter, r *http.Request) {

}

func (bh BookHandler) handlePostBook(w http.ResponseWriter, r *http.Request) {
	payload := &Book{}
	json.NewDecoder(r.Body).Decode(&payload)
	bookId, err := bh.db.Add(r.Context(), payload)
	if err != nil {
		JsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Json(w, BookIdResponse{BookID: bookId}, http.StatusOK)
}

type BookIdResponse struct {
	BookID string `json:"bookId"`
}

func (bh BookHandler) handlePatchBook(w http.ResponseWriter, r *http.Request) {

}

func (bh BookHandler) handleDeleteBook(w http.ResponseWriter, r *http.Request) {

}
