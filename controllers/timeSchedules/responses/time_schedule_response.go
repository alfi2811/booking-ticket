package responses

import (
	"booking-ticket/business/timeSchedules"
	"time"
)

type TimeScheduleResponse struct {
	ID         int       `json:"id" `
	ScheduleId int       `json:"scheduleId"`
	TimeMovie  time.Time `json:"timeMovie"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func FromDomain(domain timeSchedules.Domain) TimeScheduleResponse {
	return TimeScheduleResponse{
		ID:         domain.ID,
		ScheduleId: domain.ScheduleId,
		TimeMovie:  domain.TimeMovie,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func FromListDomain(data []timeSchedules.Domain) (result []TimeScheduleResponse) {
	result = []TimeScheduleResponse{}
	for _, time := range data {
		result = append(result, FromDomain(time))
	}
	return
}
