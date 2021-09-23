package requests

import "booking-ticket/business/cinemas"

type CinemaAdd struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Maps     string `json:"maps"`
	Phone    string `json:"phone"`
}

func (cinema *CinemaAdd) ToDomain() cinemas.Domain {
	return cinemas.Domain{
		Name:     cinema.Name,
		Location: cinema.Location,
		Maps:     cinema.Maps,
		Phone:    cinema.Phone,
	}
}
