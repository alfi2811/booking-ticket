package requests

import (
	"booking-ticket/business/movies"
	"time"
)

type MovieAdd struct {
	Title        string    `json:"title"`
	Poster       string    `json:"poster"`
	ReleaseDate  time.Time `json:"releaseDate"`
	Duration     string    `json:"duration"`
	MovieTrailer string    `json:"movieTrailer"`
	Players      string    `json:"players"`
}

func (movie *MovieAdd) ToDomain() movies.Domain {
	return movies.Domain{
		Title:        movie.Title,
		Poster:       movie.Poster,
		ReleaseDate:  movie.ReleaseDate,
		Duration:     movie.Duration,
		MovieTrailer: movie.MovieTrailer,
		Players:      movie.Players,
	}
}
