package responses

import (
	"booking-ticket/business/bookings"
	"time"
)

type BookingDetailResponse struct {
	Id          int       `json:"id"`
	TitleMovie  string    `json:"titleMovie"`
	PosterMovie string    `json:"poster"`
	NameCinema  string    `json:"nameMovie"`
	Time        time.Time `json:"time"`
	NumberSeat  string    `json:"numberSeat"`
	Quantity    int       `json:"quantity"`
	TotalPrice  int       `json:"totalPrice"`
	QrCode      string    `json:"qrcode"`
}

func FromDomainDetail(domain bookings.Domain) BookingDetailResponse {
	return BookingDetailResponse{
		Id:          domain.ID,
		TitleMovie:  domain.Detail.TitleMovie,
		PosterMovie: domain.Detail.PosterMovie,
		NameCinema:  domain.Detail.NameCinema,
		Time:        domain.Detail.Time,
		NumberSeat:  domain.NumberSeat,
		Quantity:    domain.Quantity,
		TotalPrice:  domain.TotalPrice,
		QrCode:      domain.QrCode,
	}
}
