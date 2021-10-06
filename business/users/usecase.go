package users

import (
	"booking-ticket/app/middlewares"
	"booking-ticket/business"
	"booking-ticket/helpers/encrypt"
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

	temp := encrypt.ValidateHash(password, user.Password)
	if !temp {
		return Domain{}, business.ErrPasswordWrong
	}

	user.Token, err = uc.ConfigJWT.GenerateToken(user.ID, "user")
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) Register(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" || domain.Fullname == "" || domain.Gender == "" || domain.Password == "" {
		return Domain{}, errors.New("please input all field")
	}
	var errPass error
	domain.Password, errPass = encrypt.Hash(domain.Password)
	if errPass != nil {
		return Domain{}, errPass
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

func (uc *UserUsecase) ListUserBooking(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.ListUserBooking(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}
