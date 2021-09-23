package responses

import (
	"booking-ticket/business/movies"
	"time"
)

type MovieResponse struct {
	ID          int       `json:"id" `
	Title       string    `json:"title"`
	Poster      string    `json:"Poster"`
	ReleaseDate time.Time `json:"releaseDate"`
	Status      int       `json:"status"`
}

func FromDomain(domain movies.Domain) MovieResponse {
	return MovieResponse{
		ID:          domain.ID,
		Title:       domain.Title,
		Poster:      domain.Poster,
		ReleaseDate: domain.ReleaseDate,
		Status:      domain.Status,
	}
}
