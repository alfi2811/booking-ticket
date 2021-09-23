package schedules

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	MovieId   int
	CinemaId  int
	Date      time.Time
	Price     int
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	AddSchedule(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	AddSchedule(ctx context.Context, domain Domain) (Domain, error)
}
