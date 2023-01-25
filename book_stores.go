package main

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrWrongType    = errors.New("wrong type assertion")
	ErrBookNotFound = errors.New("no book found")
	ErrNoProp       = errors.New("no prop found")
)

func NewPostgresDB(ctx context.Context, uri string) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, err
	}
	conf.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return conn.Ping(ctx)
	}
	return pgxpool.NewWithConfig(ctx, conf)
}

type PGXConn interface {
	Query(ctx context.Context, query string, params ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, params ...any) pgx.Row

	Exec(ctx context.Context, query string, params ...any) (pgconn.CommandTag, error)

	Begin(ctx context.Context) (pgx.Tx, error)
}

type BookStore struct {
	db PGXConn
}

func (bs BookStore) MigrateBookStore(ctx context.Context) error {
	authorTable := `
        CREATE TABLE IF NOT EXISTS authors (
            id SERIAl PRIMARY KEY,
            firstname TEXT,
            lastname TEXT,
            created_on TIMESTAMPTZ DEFAULT now(),
            updated_on TIMESTAMPTZ DEFAULT now(),
            deleted_on TIMESTAMPTZ
        );
    `

	bookTable := `
        CREATE TABLE IF NOT EXISTS books (
            id SERIAl PRIMARY KEY,
            title TEXT NOT NULL,
            author_id INTEGER,
            released_date date,
            created_on TIMESTAMPTZ DEFAULT now(),
            updated_on TIMESTAMPTZ DEFAULT now(),
            deleted_on TIMESTAMPTZ,

            FOREIGN KEY(author_id) REFERENCES authors(id)
        );
    `

	if _, err := bs.db.Exec(ctx, authorTable); err != nil {
		return err
	}
	if _, err := bs.db.Exec(ctx, bookTable); err != nil {
		return err
	}
	return nil
}

func NewBookStore(db PGXConn) BookStore {
	return BookStore{db: db}
}

func (bs BookStore) GetAll(ctx context.Context) ([]*Book, error) {
	return nil, nil
}

func (bs BookStore) Add(ctx context.Context, book *Book) (string, error) {
	return "", nil
}

func (bs BookStore) AddAuthor(ctx context.Context, author *Author) (string, error) {
	return "", nil
}

func (bs BookStore) GetByID(ctx context.Context, id string) (*Book, *Author, error) {
	return nil, nil, nil
}

func (bs BookStore) Delete(ctx context.Context, id string) error {
	return nil
}

func (bs BookStore) Update(ctx context.Context, id string, field BookProp, value any) (*Book, error) {
	return nil, nil
}
