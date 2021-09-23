package responses

import (
	"booking-ticket/business/schedules"
	"time"
)

type ScheduleResponse struct {
	ID        int       `json:"id" `
	MovieId   int       `json:"movieId"`
	CinemaId  int       `json:"cinemaId"`
	Date      time.Time `json:"date"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain schedules.Domain) ScheduleResponse {
	return ScheduleResponse{
		ID:        domain.ID,
		MovieId:   domain.MovieId,
		CinemaId:  domain.CinemaId,
		Date:      domain.Date,
		Price:     domain.Price,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
