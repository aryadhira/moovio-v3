package collector

import (
	"context"
	"errors"
	"fmt"
	"log"
	"moovio-v3/moovio/models"
	"moovio-v3/moovio/storages"
	"moovio-v3/utils"
	"os"
	"time"

	"github.com/google/uuid"
)

type MovieCollector struct {
	db storages.MoovioRepo
}

func NewMovieCollector(db storages.MoovioRepo) *MovieCollector {
	return &MovieCollector{
		db: db,
	}
}

func (m *MovieCollector) FetchMovieData(ctx context.Context) error {
	fetchingstart := time.Now()
	baseUrl := os.Getenv("YTS_URL")
	limit := os.Getenv("YTS_MOVIE_LIMIT")

	apiUrl := fmt.Sprintf("%v?limit=%v", baseUrl, limit)
	result, err := utils.APICall(apiUrl)
	if err != nil {
		return err
	}

	err = m.TransformAndSaveMovie(ctx, result)
	if err != nil {
		return err
	}

	log.Println("Fetching Movie Data Done at:", time.Since(fetchingstart))

	return nil
}

func (m *MovieCollector) TransformAndSaveMovie(ctx context.Context, datas map[string]interface{}) error {
	log.Println("Start Transforming Movie Data...")
	populatedate := time.Now()
	dataresult := datas["data"]
	if dataresult == nil {
		return errors.New("empty data result")
	}

	dataobj := datas["data"].(map[string]interface{})
	movies := dataobj["movies"]

	if movies == nil {
		return errors.New("empty movies data")
	}

	moviedatas := movies.([]interface{})
	transformedMovies := []models.Movies{}

	for _, each := range moviedatas {
		obj := each.(map[string]interface{})

		movie := models.Movies{}
		movie.Title = utils.InterfaceToString(obj["title"])
		movie.Year = int(utils.InterfaceToFloat64(obj["year"]))
		movie.Cover = utils.InterfaceToString(obj["large_cover_image"])
		movie.Slug = utils.InterfaceToString(obj["slug"])
		movie.Rating = utils.InterfaceToFloat64(obj["rating"])
		movie.Synopsis = utils.InterfaceToString(obj["synopsis"])
		movie.PopulateDate = populatedate
		if obj["genres"] != nil {
			movie.Categories = utils.ArrayInterfaceToString(obj["genres"].([]interface{}))
		}

		torrents := obj["torrents"].([]interface{})
		for _, torrent := range torrents {
			torrentobj := torrent.(map[string]interface{})
			movie.Id = uuid.NewString()
			movie.Quality = utils.InterfaceToString(torrentobj["quality"])
			movie.Hash = utils.InterfaceToString(torrentobj["hash"])
			magneturl := utils.GenerateMagnetUrl(movie.Title, movie.Hash, movie.Quality)
			movie.MagnetUrl = magneturl

			transformedMovies = append(transformedMovies, movie)
		}
	}

	log.Println("Start Inserting Movie Data...")
	err := m.db.InsertMoviesBulk(ctx, transformedMovies)
	if err != nil {
		return err
	}

	log.Println("Transform Movie Data Done at:", time.Since(populatedate))

	return nil
}
