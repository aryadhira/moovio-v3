package storages

import (
	"context"
	"moovio-v3/moovio/models"
)

type MigrationRepo interface {
	InitiateTable(ctx context.Context) error
	GetDBVersion(ctx context.Context) (int, error)
	InsertDBVersion(ctx context.Context, version int) error
}

type CollectorRepo interface {
	InsertMoviesBulk(ctx context.Context, movies []models.Movies) error
	GetMovieList(ctx context.Context) ([]models.Movies, error)
}

type MoovioRepo interface {
	MigrationRepo
	CollectorRepo
}
