package main

import (
	"context"
	"log"
	"moovio-v3/moovio/migration"
	"moovio-v3/moovio/services"
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

	ctx := context.Background()

	migration := migration.New(db)
	err = migration.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}

	svc := services.NewMoovioSvc(ctx, db)
	apiHandler := services.NewHandler(svc)
	err = apiHandler.Start()
	if err != nil {
		log.Fatal(err)
	}
}
