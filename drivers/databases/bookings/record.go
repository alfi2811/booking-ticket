package bookings

import (
	"booking-ticket/business/bookings"
	"time"
)

type BookingDetail struct {
	ID          int
	NumberSeat  string
	Quantity    int
	TotalPrice  int
	QrCode      string
	TitleMovie  string
	PosterMovie string
	NameCinema  string
	Time        time.Time
}

type Bookings struct {
	ID             int `gorm:"primaryKey"`
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

func (booking *Bookings) ToDomain() bookings.Domain {
	return bookings.Domain{
		ID:             booking.ID,
		UserId:         booking.UserId,
		TimeScheduleId: booking.TimeScheduleId,
		NumberSeat:     booking.NumberSeat,
		Quantity:       booking.Quantity,
		TotalPrice:     booking.TotalPrice,
		QrCode:         booking.QrCode,
		Status:         booking.Status,
		CreatedAt:      booking.CreatedAt,
		UpdatedAt:      booking.UpdatedAt,
	}
}

func (booking *BookingDetail) ToDomainDetail() bookings.Domain {
	detailBooking := bookings.DomainJoin{
		TitleMovie:  booking.TitleMovie,
		PosterMovie: booking.PosterMovie,
		NameCinema:  booking.NameCinema,
		Time:        booking.Time,
	}

	return bookings.Domain{
		ID:         booking.ID,
		NumberSeat: booking.NumberSeat,
		Quantity:   booking.Quantity,
		TotalPrice: booking.TotalPrice,
		QrCode:     booking.QrCode,
		Detail:     detailBooking,
	}
}

func ToListDomain(data []Bookings) (result []bookings.Domain) {
	result = []bookings.Domain{}
	for _, booking := range data {
		result = append(result, booking.ToDomain())
	}
	return
}

func FromDomain(domain bookings.Domain) Bookings {
	return Bookings{
		UserId:         domain.UserId,
		TimeScheduleId: domain.TimeScheduleId,
		NumberSeat:     domain.NumberSeat,
		Quantity:       domain.Quantity,
		TotalPrice:     domain.TotalPrice,
		QrCode:         domain.QrCode,
		Status:         1,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}
