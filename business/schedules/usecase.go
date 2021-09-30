package schedules

import (
	"context"
	"errors"
	"time"
)

type ScheduleUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewScheduleUsecase(repo Repository, timeout time.Duration) Usecase {
	return &ScheduleUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *ScheduleUsecase) AddSchedule(ctx context.Context, domain Domain) (Domain, error) {
	if domain.MovieId == 0 || domain.CinemaId == 0 || domain.Price == 0 {
		return Domain{}, errors.New("please input all field")
	}
	user, err := uc.Repo.AddSchedule(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (sc *ScheduleUsecase) ListSchedule(ctx context.Context) ([]Domain, error) {
	schedules, err := sc.Repo.ListSchedule(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return schedules, nil
}

func (sc *ScheduleUsecase) DetailTimeSchedule(ctx context.Context, idSchedule int) (Domain, error) {
	schedule, err := sc.Repo.DetailTimeSchedule(ctx, idSchedule)

	if err != nil {
		return Domain{}, err
	}
	return schedule, nil
}
