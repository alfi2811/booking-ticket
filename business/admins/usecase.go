package admins

import (
	"booking-ticket/app/middlewares"
	"booking-ticket/helpers/encrypt"
	"context"
	"errors"
	"time"
)

type AdminUsecase struct {
	ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewAdminUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &AdminUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

// core bisinis login
func (uc *AdminUsecase) Login(ctx context.Context, email string, password string) (Domain, error) {
	if email == "" {
		return Domain{}, errors.New("email empty")
	}
	if password == "" {
		return Domain{}, errors.New("password empty")
	}

	user, err := uc.Repo.Login(ctx, email, password)
	// temp := encrypt.ValidateHash(password, user.Password)
	// if !temp {
	// 	return Domain{}, errors.New("password wrong")
	// }

	if err != nil {
		return Domain{}, err
	}
	user.Token, err = uc.ConfigJWT.GenerateToken(user.ID, "admin")
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *AdminUsecase) Register(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" || domain.Password == "" {
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

func (uc *AdminUsecase) GetAdmin(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetAdmin(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}
