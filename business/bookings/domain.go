package bookings

import (
	"context"
	"time"
)

type DomainJoin struct {
	ID          int
	TitleMovie  string
	PosterMovie string
	NameCinema  string
	Time        time.Time
}

type Domain struct {
	ID             int
	UserId         int
	TimeScheduleId int
	NumberSeat     string
	Quantity       int
	TotalPrice     int
	QrCode         string
	Status         int
	TitleMovie     string
	PosterMovie    string
	NameCinema     string
	Time           time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Usecase interface {
	AddBooking(ctx context.Context, domain Domain) (Domain, error)
	ListBooking(ctx context.Context) ([]Domain, error)
	ListBookingUser(ctx context.Context, userId int) ([]DomainJoin, error)
}

type Repository interface {
	AddBooking(ctx context.Context, domain Domain) (Domain, error)
	ListBooking(ctx context.Context) ([]Domain, error)
	ListBookingUser(ctx context.Context, userId int) ([]DomainJoin, error)
}
