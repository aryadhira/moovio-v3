package services

import (
	"context"
	"moovio-v3/moovio/models"
	"moovio-v3/moovio/storages"
)

type MoovioSvc struct {
	db  storages.MoovioRepo
	ctx context.Context
}

func NewMoovioSvc(ctx context.Context, db storages.MoovioRepo) *MoovioSvc {
	return &MoovioSvc{
		db:  db,
		ctx: ctx,
	}
}

func (m *MoovioSvc) GetMovieList() ([]models.Movies, error) {
	res, err := m.db.GetMovieList(m.ctx)
	if err != nil {
		return res, err
	}

	return res, nil
}
