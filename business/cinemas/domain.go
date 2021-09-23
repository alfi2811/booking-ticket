package cinemas

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Location  string
	Maps      string
	Phone     string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	AddCinema(ctx context.Context, domain Domain) (Domain, error)
	ListCinema(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	AddCinema(ctx context.Context, domain Domain) (Domain, error)
	ListCinema(ctx context.Context) ([]Domain, error)
}
