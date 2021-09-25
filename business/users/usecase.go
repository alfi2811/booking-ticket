package users

import (
	"booking-ticket/app/middlewares"
	"context"
	"errors"
	"time"
)

type UserUsecase struct {
	ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

// core bisinis login
func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (Domain, error) {
	if email == "" {
		return Domain{}, errors.New("email empty")
	}
	if password == "" {
		return Domain{}, errors.New("password empty")
	}

	user, err := uc.Repo.Login(ctx, email, password)

	if err != nil {
		return Domain{}, err
	}
	user.Token, err = uc.ConfigJWT.GenerateToken(user.ID)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) Register(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	user, err := uc.Repo.Register(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) GetUser(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetUser(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}
