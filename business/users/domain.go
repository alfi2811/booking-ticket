package users

import (
	"booking-ticket/controllers/users/requests"
	"context"
	"time"
)

type Domain struct {
	ID        int
	Email     string
	Password  string
	Fullname  string
	Gender    string
	Phone     string
	Status    int
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	Register(ctx context.Context, userRegister requests.UserRegister) (Domain, error)
	GetUser(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	Register(ctx context.Context, userRegister requests.UserRegister) (Domain, error)
	GetUser(ctx context.Context) ([]Domain, error)
}
