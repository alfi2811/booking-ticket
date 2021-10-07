package responses

import (
	"booking-ticket/business/schedules"
	"booking-ticket/controllers/timeSchedules/responses"
	"time"
)

type ScheduleMovie struct {
	Title    string `json:"title"`
	Poster   string `json:"poster"`
	Duration string `json:"duration"`
}

type ScheduleCinema struct {
	Name string `json:"name"`
	Maps string `json:"maps"`
}

type ScheduleListResponse struct {
	ID        int            `json:"id" `
	Movie     ScheduleMovie  `json:"movie"`
	Cinema    ScheduleCinema `json:"cinema"`
	Date      time.Time      `json:"date"`
	Price     int            `json:"price"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

type ScheduleDetail struct {
	Schedule ScheduleListResponse             `json:"schedule"`
	Time     []responses.TimeScheduleResponse `json:"time"`
}

func FromDomainJoin(domain schedules.Domain) ScheduleListResponse {
	scheduleMovie := ScheduleMovie{
		Title:    domain.Movie.Title,
		Poster:   domain.Movie.Poster,
		Duration: domain.Movie.Duration,
	}

	scheduleCinema := ScheduleCinema{
		Name: domain.Cinema.Name,
		Maps: domain.Cinema.Maps,
	}
	return ScheduleListResponse{
		ID:        domain.ID,
		Movie:     scheduleMovie,
		Cinema:    scheduleCinema,
		Date:      domain.Date,
		Price:     domain.Price,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDetailJoin(domain schedules.Domain) ScheduleDetail {
	return ScheduleDetail{
		Schedule: FromDomainJoin(domain),
		Time:     responses.FromListDomain(domain.Time),
	}
}

func FromListDomainJoin(data []schedules.Domain) (result []ScheduleListResponse) {
	result = []ScheduleListResponse{}
	for _, schedule := range data {
		result = append(result, FromDomainJoin(schedule))
	}
	return
}
