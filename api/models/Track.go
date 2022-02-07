package models

type Track struct {
	Id         string
	Name       string
	Popularity float64
	Acousticness float64
	Danceability float64
	Duration_ms float64
	Energy float64
	Instrumentalness float64
	Liveness float64
	Loudness float64
	Mode float64
	Speechiness float64
	Tempo float64
	Valence float64
}