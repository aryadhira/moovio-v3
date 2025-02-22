package postgres

import (
	"context"
	"fmt"
	"moovio-v3/moovio/models"
	"strings"
)

func (p *PostgresRepo) InsertMoviesBulk(ctx context.Context, movies []models.Movies) error {
	query := `
		INSERT INTO movie_list(
		id,
		title,
		year,
		synopsis,
		rating,
		cover,
		quality,
		hash,
		slug,
		magnet_url,
		categories,
		populate_date
		)
		VALUES 
	`
	queryValStr := generateQueryString(movies)

	finalQuery := query + queryValStr
	_, err := p.Db.ExecContext(ctx, finalQuery)
	if err != nil {
		return err
	}

	return nil
}

func generateQueryString(movies []models.Movies) string {
	queryStr := ""
	for i, each := range movies {
		query := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v",
			`'`+each.Id+`'`,
			`'`+strings.ReplaceAll(each.Title, "'", "")+`'`,
			each.Year,
			`'`+strings.ReplaceAll(each.Synopsis, "'", "")+`'`,
			each.Rating,
			`'`+each.Cover+`'`,
			`'`+each.Quality+`'`,
			`'`+each.Hash+`'`,
			`'`+each.Slug+`'`,
			`'`+each.MagnetUrl+`'`,
			`'`+each.Categories+`'`,
			"NOW()::TIMESTAMP",
		)
		if i == len(movies)-1 {
			queryStr += "(" + query + ")"
		} else {
			queryStr += "(" + query + "),"
		}
	}
	return queryStr
}

func (p *PostgresRepo) GetMovieList(ctx context.Context) ([]models.Movies, error) {
	query := `
		SELECT 
			title,
			year,
			rating,
			cover
		FROM movie_list
		GROUP BY title, year, rating, cover
	`
	rows, err := p.Db.QueryContext(ctx, query)
	if err != nil {
		return []models.Movies{}, err
	}
	defer rows.Close()

	result := []models.Movies{}
	for rows.Next() {
		movie := models.Movies{}
		err = rows.Scan(&movie.Title, &movie.Year, &movie.Rating, &movie.Cover)
		if err != nil {
			return result, err
		}
		result = append(result, movie)
	}

	return result, nil
}
