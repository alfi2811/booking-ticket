package responses

import (
	"booking-ticket/business/bookings"
	"time"
)

type BookingUserResponse struct {
	Id          int       `json:"id"`
	TitleMovie  string    `json:"titleMovie"`
	PosterMovie string    `json:"posterMovie"`
	NameCinema  string    `json:"nameCinema"`
	Time        time.Time `json:"time"`
}

func FromDomainBookingUser(domain bookings.DomainJoin) BookingUserResponse {
	return BookingUserResponse{
		Id:          domain.ID,
		TitleMovie:  domain.TitleMovie,
		PosterMovie: domain.PosterMovie,
		NameCinema:  domain.NameCinema,
		Time:        domain.Time,
	}
}

func FromListDomainBookingUser(data []bookings.DomainJoin) (result []BookingUserResponse) {
	result = []BookingUserResponse{}
	for _, movie := range data {
		result = append(result, FromDomainBookingUser(movie))
	}
	return
}
