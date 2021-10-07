package cinemas

import (
	"context"
	"errors"
	"time"
)

type CinemaUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewCinemaUsecase(repo Repository, timeout time.Duration) Usecase {
	return &CinemaUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *CinemaUsecase) AddCinema(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" || domain.Location == "" || domain.Maps == "" || domain.Phone == "" {
		return Domain{}, errors.New("please input all field")
	}
	user, err := uc.Repo.AddCinema(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (cc *CinemaUsecase) ListCinema(ctx context.Context) ([]Domain, error) {
	movie, err := cc.Repo.ListCinema(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return movie, nil
}

func (cc *CinemaUsecase) CinemaDetail(ctx context.Context, cinemaId int) (Domain, error) {
	movie, err := cc.Repo.CinemaDetail(ctx, cinemaId)

	if err != nil {
		return Domain{}, err
	}
	return movie, nil
}
