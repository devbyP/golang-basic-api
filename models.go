package main

import (
	"time"
)

type BookProp int

const (
	BookTitle BookProp = iota + 1
	BookAuthorID
	BookTotalPage
	BookReleasedDate
)

type Book struct {
	ID           string     `json:"id"`
	Title        string     `json:"title"`
	AuthorID     string     `json:"authorId"`
	TotalPage    int        `json:"totalPage"`
	ReleasedDate *time.Time `json:"releasedDate"`
	CreatedOn    time.Time  `json:"createdOn"`
	UpdatedOn    time.Time  `json:"updatedOn"`
	DeletedOn    *time.Time `json:"deletedOn"`
}

type Author struct {
	ID        string     `json:"id"`
	Firstname string     `json:"firstname"`
	Lastname  string     `json:"lastname"`
	CreatedOn time.Time  `json:"createOn"`
	UpdatedOn time.Time  `json:"updatedOn"`
	DeletedOn *time.Time `json:"deletedOn"`
}
