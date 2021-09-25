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
	Status         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Usecase interface {
	AddBooking(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	AddBooking(ctx context.Context, domain Domain) (Domain, error)
}
