package schedules

import (
	"booking-ticket/business/schedules"
	"time"
)

type Schedules struct {
	ID        int `gorm:"primaryKey"`
	MovieId   int
	CinemaId  int
	Date      time.Time
	Price     int
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (schedule *Schedules) ToDomain() schedules.Domain {
	return schedules.Domain{
		ID:        schedule.ID,
		MovieId:   schedule.MovieId,
		CinemaId:  schedule.CinemaId,
		Date:      schedule.Date,
		Price:     schedule.Price,
		Status:    schedule.Status,
		CreatedAt: schedule.CreatedAt,
		UpdatedAt: schedule.UpdatedAt,
	}
}

func FromDomain(domain schedules.Domain) Schedules {
	return Schedules{
		MovieId:  domain.MovieId,
		CinemaId: domain.CinemaId,
		Date:     domain.Date,
		Price:    domain.Price,
		Status:   1,
	}
}
