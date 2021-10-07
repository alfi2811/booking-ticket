package users

import (
	"booking-ticket/drivers/databases/bookings"
	"context"
	"time"
)

type Domain struct {
	ID        int
	Email     string
	Password  string
	Fullname  string
	Gender    string
	Phone     string
	Status    int
	Token     string
	Booking   []bookings.Bookings
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	Register(ctx context.Context, domain Domain) (Domain, error)
	GetUser(ctx context.Context) ([]Domain, error)
	ListUserBooking(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	Register(ctx context.Context, domain Domain) (Domain, error)
	GetUser(ctx context.Context) ([]Domain, error)
	ListUserBooking(ctx context.Context) ([]Domain, error)
}
