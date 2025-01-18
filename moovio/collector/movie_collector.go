package collector

import (
	"moovio-v3/moovio/storages"
)

type MovieCollector struct {
	db storages.MoovioRepo
}

func NewMovieCollector(db storages.MoovioRepo) *MovieCollector {
	return &MovieCollector{
		db: db,
	}
}
