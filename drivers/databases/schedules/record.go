package schedules

import (
	"booking-ticket/business/cinemas"
	"booking-ticket/business/movies"

	// movieDB "booking-ticket/drivers/databases/movies"
	"booking-ticket/business/schedules"
	"booking-ticket/drivers/databases/timeSchedules"
	"time"
)

type Schedules struct {
	ID           int `gorm:"primaryKey"`
	MovieId      int
	CinemaId     int
	Date         time.Time
	Price        int
	Status       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	TimeSchedule []timeSchedules.TimeSchedules `gorm:"foreignKey:ScheduleId;references:ID"`
}

type ListSchedules struct {
	ID       int
	Title    string
	Poster   string
	Duration string
	Name     string
	Maps     string
	Date     time.Time
	Price    int
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

func (schedule *ListSchedules) ToDomainJoin() schedules.Domain {
	Movie := movies.Domain{
		Title:    schedule.Title,
		Poster:   schedule.Poster,
		Duration: schedule.Duration,
	}

	Cinema := cinemas.Domain{
		Name: schedule.Name,
		Maps: schedule.Maps,
	}

	return schedules.Domain{
		ID:     schedule.ID,
		Movie:  Movie,
		Cinema: Cinema,
		Date:   schedule.Date,
		Price:  schedule.Price,
	}
}

func ToListDomain(data []Schedules) (result []schedules.Domain) {
	result = []schedules.Domain{}
	for _, schedule := range data {
		result = append(result, schedule.ToDomain())
	}
	return
}

func ToListDomainJoin(data []ListSchedules) (result []schedules.Domain) {
	result = []schedules.Domain{}
	for _, schedule := range data {
		result = append(result, schedule.ToDomainJoin())
	}
	return
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
