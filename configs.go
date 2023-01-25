package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	port  string
	dbURI string
)

func setupEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("cannot load env: %v\n", err)
	}

	port = os.Getenv("PORT")
	if port == "" {
		// default port to 3000.
		port = "3000"
	}

	err := setdb()
	if err != nil {
		log.Fatalln(err)
	}
}

func setdb() error {
	dbURI = os.Getenv("DB_URI")
	if dbURI == "" {
		// no default db
		return fmt.Errorf("default db not allowed")
	}
	return nil
}
