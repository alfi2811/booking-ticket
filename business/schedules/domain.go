package schedules

import (
	"booking-ticket/business/cinemas"
	"booking-ticket/business/movies"
	"booking-ticket/business/timeSchedules"
	"context"
	"time"
)

type Domain struct {
	ID        int
	MovieId   int
	CinemaId  int
	Title     string
	Movie     movies.Domain
	Cinema    cinemas.Domain
	Time      []timeSchedules.Domain
	Date      time.Time
	Price     int
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	AddSchedule(ctx context.Context, domain Domain) (Domain, error)
	ListSchedule(ctx context.Context) ([]Domain, error)
	DetailTimeSchedule(ctx context.Context, idSchedule int) (Domain, error)
}

type Repository interface {
	AddSchedule(ctx context.Context, domain Domain) (Domain, error)
	ListSchedule(ctx context.Context) ([]Domain, error)
	DetailTimeSchedule(ctx context.Context, idSchedule int) (Domain, error)
}
