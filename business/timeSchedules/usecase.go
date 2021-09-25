package timeSchedules

import (
	"context"
	"errors"
	"time"
)

type TimeScheduleUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewTimeScheduleUsecase(repo Repository, timeout time.Duration) Usecase {
	return &TimeScheduleUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *TimeScheduleUsecase) AddScheduleTime(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ScheduleId == 0 {
		return Domain{}, errors.New("please input all field")
	}
	user, err := uc.Repo.AddScheduleTime(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *TimeScheduleUsecase) DetailScheduleTime(ctx context.Context, idSchedule int) ([]Domain, error) {
	if idSchedule == 0 {
		return []Domain{}, errors.New("please input the id")
	}
	schedule, err := uc.Repo.DetailScheduleTime(ctx, idSchedule)

	if err != nil {
		return []Domain{}, err
	}
	return schedule, nil
}
