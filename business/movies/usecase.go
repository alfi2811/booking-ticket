package movies

import (
	"context"
	"errors"
	"time"
)

type MovieUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewMovieUsecase(repo Repository, timeout time.Duration) Usecase {
	return &MovieUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *MovieUsecase) AddMovie(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Title == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Duration == "" {
		return Domain{}, errors.New("password empty")
	}
	movie, err := uc.Repo.AddMovie(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return movie, nil
}

func (mc *MovieUsecase) ListMovie(ctx context.Context) ([]Domain, error) {
	movie, err := mc.Repo.ListMovie(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return movie, nil
}
