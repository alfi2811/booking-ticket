package schedules

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	MovieId   int
	Title     string
	CinemaId  int
	Date      time.Time
	Price     int
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	AddSchedule(ctx context.Context, domain Domain) (Domain, error)
	ListSchedule(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	AddSchedule(ctx context.Context, domain Domain) (Domain, error)
	ListSchedule(ctx context.Context) ([]Domain, error)
}
