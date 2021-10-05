package bookings

import (
	"booking-ticket/business/qrcode"
	"context"
	"errors"
	"strconv"
	"time"
)

type BookingUsecase struct {
	Repo           Repository
	QrCodeRepo     qrcode.Repository
	contextTimeout time.Duration
}

func NewBookingUsecase(repo Repository, qrCodeRepo qrcode.Repository, timeout time.Duration) Usecase {
	return &BookingUsecase{
		Repo:           repo,
		QrCodeRepo:     qrCodeRepo,
		contextTimeout: timeout,
	}
}

func (uc *BookingUsecase) AddBooking(ctx context.Context, domain Domain) (Domain, error) {
	if domain.UserId == 0 || domain.TimeScheduleId == 0 {
		return Domain{}, errors.New("please input user or time schedule id")
	}
	if domain.NumberSeat == "" || domain.Quantity == 0 || domain.TotalPrice == 0 {
		return Domain{}, errors.New("please input all field ")
	}
	dataUser := "{userId:" + strconv.Itoa(domain.UserId) + ",timeScheduleId:" + strconv.Itoa(domain.TimeScheduleId) + "}"
	dataQr, errQr := uc.QrCodeRepo.GetQrCode(ctx, dataUser)
	if errQr != nil {
		return Domain{}, errQr
	}
	domain.QrCode = dataQr.QrCode
	user, err := uc.Repo.AddBooking(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
func (bc *BookingUsecase) ListBooking(ctx context.Context) ([]Domain, error) {
	booking, err := bc.Repo.ListBooking(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return booking, nil
}

func (bc *BookingUsecase) ListBookingUser(ctx context.Context, userId int) ([]DomainJoin, error) {
	booking, err := bc.Repo.ListBookingUser(ctx, userId)

	if err != nil {
		return []DomainJoin{}, err
	}
	return booking, nil
}

func (bc *BookingUsecase) DetailBooking(ctx context.Context, bookingId int) (Domain, error) {
	booking, err := bc.Repo.DetailBooking(ctx, bookingId)

	if err != nil {
		return Domain{}, err
	}
	return booking, nil
}
