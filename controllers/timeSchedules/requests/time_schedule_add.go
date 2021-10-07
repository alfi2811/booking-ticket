package requests

import (
	"booking-ticket/business/timeSchedules"
	"time"
)

type TimeScheduleAdd struct {
	ScheduleId int       `json:"scheduleId"`
	TimeMovie  time.Time `json:"timeMovie"`
}

func (timescheduleAdd *TimeScheduleAdd) ToDomain() timeSchedules.Domain {
	return timeSchedules.Domain{
		ScheduleId: timescheduleAdd.ScheduleId,
		TimeMovie:  timescheduleAdd.TimeMovie,
	}
}
