package postgres

import (
	"context"

	"github.com/google/uuid"
)

func (p *PostgresRepo) InitiateTable(ctx context.Context) error {
	query := `
		CREATE TABLE IF NOT EXISTS db_version(
			id UUID PRIMARY KEY,
			version INT
		);
		
		CREATE TABLE IF NOT EXISTS movie_list(
			id UUID PRIMARY KEY,
			title TEXT,
			year INT,
			synopsis TEXT,
			rating DECIMAL,
			cover TEXT,
			quality TEXT,
			hash TEXT,
			slug TEXT,
			magnet_url TEXT,
			categories TEXT,
			populate_date TIMESTAMP
		);


	`
	_, err := p.Db.QueryContext(ctx, query)
	if err != nil {
		return err
	}

	err = p.InsertDBVersion(ctx, 1)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepo) GetDBVersion(ctx context.Context) (int, error) {
	query := `
		SELECT version 
		FROM db_version
		ORDER BY version DESC
		LIMIT 1
	`
	rows, err := p.Db.QueryContext(ctx, query)
	if err != nil {
		return 0, err
	}

	defer rows.Close()
	dbversion := 0
	for rows.Next() {
		rows.Scan(&dbversion)
	}

	return dbversion, nil
}

func (p *PostgresRepo) InsertDBVersion(ctx context.Context, version int) error {
	query := `
		INSERT INTO db_version (id, version)
		VALUES ($1,$2);
	`

	statement, err := p.Db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, uuid.NewString(), version)
	if err != nil {
		return err
	}

	return nil
}
