package responses

import (
	"booking-ticket/business/movies"
	"time"
)

type MovieResponse struct {
	ID          int       `json:"id" `
	Title       string    `json:"title"`
	Poster      string    `json:"poster"`
	ReleaseDate time.Time `json:"releaseDate"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func FromDomain(domain movies.Domain) MovieResponse {
	return MovieResponse{
		ID:          domain.ID,
		Title:       domain.Title,
		Poster:      domain.Poster,
		ReleaseDate: domain.ReleaseDate,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromListDomain(data []movies.Domain) (result []MovieResponse) {
	result = []MovieResponse{}
	for _, movie := range data {
		result = append(result, FromDomain(movie))
	}
	return
}
