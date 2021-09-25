package bookings

import (
	"context"
	"errors"
	"time"
)

type BookingUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewBookingUsecase(repo Repository, timeout time.Duration) Usecase {
	return &BookingUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *BookingUsecase) AddBooking(ctx context.Context, domain Domain) (Domain, error) {
	if domain.NumberSeat == "" || domain.Quantity == 0 || domain.TotalPrice == 0 {
		return Domain{}, errors.New("please input all field")
	}
	user, err := uc.Repo.AddBooking(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
