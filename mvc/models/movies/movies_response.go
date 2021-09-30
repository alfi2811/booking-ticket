package movies

import "time"

type MovieResponse struct {
	ID          int       `json:"id" `
	Title       string    `json:"title"`
	Poster      string    `json:"Poster"`
	ReleaseDate time.Time `json:"releaseDate"`
	Status      int       `json:"status"`
}
