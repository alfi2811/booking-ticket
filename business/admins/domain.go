package admins

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Email     string
	Password  string
	Fullname  string
	Status    int
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	Register(ctx context.Context, domain Domain) (Domain, error)
	GetAdmin(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	Register(ctx context.Context, domain Domain) (Domain, error)
	GetAdmin(ctx context.Context) ([]Domain, error)
}
