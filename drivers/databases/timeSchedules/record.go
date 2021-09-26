package timeSchedules

import (
	"booking-ticket/business/timeSchedules"
	"booking-ticket/drivers/databases/bookings"
	"time"
)

type TimeSchedules struct {
	ID         int `gorm:"primaryKey"`
	ScheduleId int
	TimeMovie  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Booking    []bookings.Bookings `gorm:"foreignKey:TimeScheduleId;references:ID"`
}

func (schedule *TimeSchedules) ToDomain() timeSchedules.Domain {
	return timeSchedules.Domain{
		ID:         schedule.ID,
		ScheduleId: schedule.ScheduleId,
		TimeMovie:  schedule.TimeMovie,
		CreatedAt:  schedule.CreatedAt,
		UpdatedAt:  schedule.UpdatedAt,
	}
}

func ToListDomain(data []TimeSchedules) (result []timeSchedules.Domain) {
	result = []timeSchedules.Domain{}
	for _, time := range data {
		result = append(result, time.ToDomain())
	}
	return
}

func FromDomain(domain timeSchedules.Domain) TimeSchedules {
	return TimeSchedules{
		ScheduleId: domain.ScheduleId,
		TimeMovie:  domain.TimeMovie,
	}
}
