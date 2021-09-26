package timeSchedules

import (
	"context"
	"time"
)

type Domain struct {
	ID         int
	ScheduleId int
	TimeMovie  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Usecase interface {
	AddScheduleTime(ctx context.Context, domain Domain) (Domain, error)
	DetailScheduleTime(ctx context.Context, idSchedule int) ([]Domain, error)
}

type Repository interface {
	AddScheduleTime(ctx context.Context, domain Domain) (Domain, error)
	DetailScheduleTime(ctx context.Context, idSchedule int) ([]Domain, error)
}
