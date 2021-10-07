package requests

import (
	"booking-ticket/business/schedules"
	"time"
)

type ScheduleAdd struct {
	MovieId  int       `json:"movieId"`
	CinemaId int       `json:"cinemaId"`
	Date     time.Time `json:"date"`
	Price    int       `json:"price"`
}

func (scheduleAdd *ScheduleAdd) ToDomain() schedules.Domain {
	return schedules.Domain{
		MovieId:  scheduleAdd.MovieId,
		CinemaId: scheduleAdd.CinemaId,
		Date:     scheduleAdd.Date,
		Price:    scheduleAdd.Price,
	}
}
