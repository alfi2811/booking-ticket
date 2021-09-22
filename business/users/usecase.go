package users

import (
	"booking-ticket/controllers/users/requests"
	"context"
	"errors"
	"time"
)

type UserUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
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
	return user, nil
}

func (uc *UserUsecase) Register(ctx context.Context, userRegister requests.UserRegister) (Domain, error) {
	if userRegister.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if userRegister.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	user, err := uc.Repo.Register(ctx, userRegister)

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
