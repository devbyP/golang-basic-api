package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	setupEnv()
	ctx := context.Background()
	// connect to database
	db, err := NewPostgresDB(ctx, dbURI)
	if err != nil {
		log.Fatalf("fail to connect to database: %v", err)
	}

	// init resuorce use in this route.
	bs := NewBookStore(db)
	bh := NewBookHandler(bs)

	err = bs.MigrateBookStore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	r := NewServer()
	r.MountHandlers(bh)

	log.Println("server run on port:" + port)
	log.Fatalln(http.ListenAndServe(":"+port, r.Router))
}
