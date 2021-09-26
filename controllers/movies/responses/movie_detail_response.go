package responses

import (
	"booking-ticket/business/movies"
	"time"
)

type MovieDetailResponse struct {
	ID           int       `json:"id" `
	Title        string    `json:"title"`
	Poster       string    `json:"poster"`
	Duration     string    `json:"duration"`
	MovieTrailer string    `json:"movieTrailer"`
	Players      string    `json:"players"`
	ReleaseDate  time.Time `json:"releaseDate"`
}

func FromDomainDetail(domain movies.Domain) MovieDetailResponse {
	return MovieDetailResponse{
		ID:           domain.ID,
		Title:        domain.Title,
		Poster:       domain.Poster,
		Duration:     domain.Duration,
		MovieTrailer: domain.MovieTrailer,
		Players:      domain.Players,
		ReleaseDate:  domain.ReleaseDate,
	}
}
