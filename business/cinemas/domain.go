package cinemas

import (
	"context"
	"time"
)

type CinemaDetail struct {
	Id       string    `json:"scheduleId"`
	IdMovie  string    `json:"movieId"`
	Title    string    `json:"title"`
	Poster   string    `json:"poster"`
	Duration string    `json:"duration"`
	Date     time.Time `json:"date"`
	Price    int       `json:"price"`
}

type Domain struct {
	ID           int
	Name         string
	Location     string
	Maps         string
	Phone        string
	Status       int
	DetailCinema []CinemaDetail
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usecase interface {
	AddCinema(ctx context.Context, domain Domain) (Domain, error)
	ListCinema(ctx context.Context) ([]Domain, error)
	CinemaDetail(ctx context.Context) (Domain, error)
}

type Repository interface {
	AddCinema(ctx context.Context, domain Domain) (Domain, error)
	ListCinema(ctx context.Context) ([]Domain, error)
	CinemaDetail(ctx context.Context) (Domain, error)
}
