package bookings

import (
	"context"
	"time"
)

type Domain struct {
	ID             int
	UserId         int
	TimeScheduleId int
	NumberSeat     string
	Quantity       int
	TotalPrice     int
	QrCode         string
	Status         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Usecase interface {
	AddBooking(ctx context.Context, domain Domain) (Domain, error)
	ListBooking(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	AddBooking(ctx context.Context, domain Domain) (Domain, error)
	ListBooking(ctx context.Context) ([]Domain, error)
}
