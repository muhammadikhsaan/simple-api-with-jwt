package config

import (
	"context"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

var (
	DB *pg.DB
)

func init() {
	ctx := context.Background()

	DB = pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_ADDRESS"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})

	if err := DB.Ping(ctx); err != nil {
		log.Fatalf("Error Connect To %s", os.Getenv("DB_NAME"))
	}

	log.Printf("Successfuly Connect to %s", os.Getenv("DB_NAME"))
}
