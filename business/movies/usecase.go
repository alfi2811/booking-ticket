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
	if domain.Title == "" || domain.Poster == "" || domain.Duration == "" || domain.Players == "" || domain.MovieTrailer == "" {
		return Domain{}, errors.New("please input all field ")
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

func (mc *MovieUsecase) DetailMovie(ctx context.Context, idMovie int) (Domain, error) {
	movie, err := mc.Repo.DetailMovie(ctx, idMovie)

	if err != nil {
		return Domain{}, err
	}
	return movie, nil
}
