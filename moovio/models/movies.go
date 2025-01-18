package models

import "time"

type Movies struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Year         int       `json:"year"`
	Synopsis     string    `json:"synopsis"`
	Rating       float64   `json:"rating"`
	Cover        string    `json:"cover"`
	Quality      string    `json:"quality"`
	Hash         string    `json:"hash"`
	Slug         string    `json:"slug"`
	MagnetUrl    string    `json:"magnet_url"`
	Categories   string    `json:"categories"`
	PopulateDate time.Time `json:"populate_date"`
}

func (m *Movies) TableName() string {
	return "movie_list"
}
