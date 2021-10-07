package movies

import (
	"context"
	"time"
)

type Domain struct {
	ID           int
	Title        string
	Poster       string
	ReleaseDate  time.Time
	Duration     string
	MovieTrailer string
	Players      string
	Status       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usecase interface {
	AddMovie(ctx context.Context, domain Domain) (Domain, error)
	ListMovie(ctx context.Context) ([]Domain, error)
	DetailMovie(ctx context.Context, idMovie int) (Domain, error)
}

type Repository interface {
	AddMovie(ctx context.Context, domain Domain) (Domain, error)
	ListMovie(ctx context.Context) ([]Domain, error)
	DetailMovie(ctx context.Context, idMovie int) (Domain, error)
}
