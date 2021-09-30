package bookings

import (
	"context"
	"errors"
	"strconv"
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
	domain.QrCode = "https://api.qrserver.com/v1/create-qr-code/?data={UserId:" + strconv.Itoa(domain.UserId) + ",TimeScheduleId:" + strconv.Itoa(domain.TimeScheduleId) + "}&size=100x100"
	user, err := uc.Repo.AddBooking(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
func (bc *BookingUsecase) ListBooking(ctx context.Context) ([]Domain, error) {
	movie, err := bc.Repo.ListBooking(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return movie, nil
}

func (bc *BookingUsecase) ListBookingUser(ctx context.Context, userId int) ([]DomainJoin, error) {
	movie, err := bc.Repo.ListBookingUser(ctx, userId)

	if err != nil {
		return []DomainJoin{}, err
	}
	return movie, nil
}

func (bc *BookingUsecase) DetailBooking(ctx context.Context, bookingId int) (Domain, error) {
	movie, err := bc.Repo.DetailBooking(ctx, bookingId)

	if err != nil {
		return Domain{}, err
	}
	return movie, nil
}
