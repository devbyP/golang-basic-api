// memory storage for testing
package main

import (
	"context"
	"time"
)

type MemDB struct {
	aId     string
	bId     string
	books   map[string]*Book
	authors map[string]*Author
}

func NewMemDB(aId, bId string) *MemDB {
	return &MemDB{
		aId:     aId,
		bId:     bId,
		books:   make(map[string]*Book),
		authors: make(map[string]*Author),
	}
}

func (m MemDB) GetAll(ctx context.Context) ([]*Book, error) {
	books := make([]*Book, 0)
	for _, b := range m.books {
		books = append(books, b)
	}
	return books, nil
}

func (m MemDB) Add(ctx context.Context, book *Book) (string, error) {
	id := m.bId
	book.ID = id
	m.books[id] = book
	return id, nil
}

func (m MemDB) AddAuthor(ctx context.Context, author *Author) (string, error) {
	id := m.aId
	author.ID = id
	m.authors[id] = author
	return id, nil
}

func (m MemDB) GetByID(ctx context.Context, id string) (*Book, *Author, error) {
	book, ok := m.books[id]
	if !ok {
		return nil, nil, ErrBookNotFound
	}
	author := m.authors[book.AuthorID]
	return book, author, nil
}

func (m MemDB) Delete(ctx context.Context, id string) error {
	if _, ok := m.books[id]; ok {
		delete(m.books, id)
	}
	return nil
}

func (m MemDB) Update(ctx context.Context, id string, field BookProp, value any) (*Book, error) {
	book, ok := m.books[id]
	if !ok {
		return nil, ErrBookNotFound
	}
	switch field {
	case BookTitle:
		newTitle, ok := value.(string)
		if !ok {
			return nil, ErrWrongType
		}
		m.books[id].Title = newTitle
	case BookAuthorID:
		newAuthorID, ok := value.(string)
		if !ok {
			return nil, ErrWrongType
		}
		m.books[id].AuthorID = newAuthorID
	case BookTotalPage:
		newTPage, ok := value.(int)
		if !ok {
			return nil, ErrWrongType
		}
		m.books[id].TotalPage = newTPage
	case BookReleasedDate:
		newRD, ok := value.(*time.Time)
		if !ok {
			return nil, ErrWrongType
		}
		m.books[id].ReleasedDate = newRD
	default:
		return nil, ErrNoProp
	}
	return book, nil
}
