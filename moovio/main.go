package main

import (
	"context"
	"log"
	"moovio-v3/moovio/migration"
	"moovio-v3/moovio/storages/postgres"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.New()
	if err != nil {
		log.Fatal(err)
	}

	migration := migration.New(db)
	err = migration.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// coll := collector.NewMovieCollector(db)
	// err = coll.CollectMovie(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
