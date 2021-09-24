package users

import (
	"booking-ticket/business/users"
	"time"
)

type Users struct {
	ID        int    `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Password  string
	Fullname  string
	Gender    string
	Phone     string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *Users) ToDomain() users.Domain {
	return users.Domain{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		Fullname:  user.Fullname,
		Gender:    user.Gender,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
func ToListDomain(data []Users) (result []users.Domain) {
	result = []users.Domain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}

func FromDomain(domain users.Domain) Users {
	return Users{
		ID:        domain.ID,
		Email:     domain.Email,
		Password:  domain.Password,
		Fullname:  domain.Fullname,
		Gender:    domain.Gender,
		Phone:     domain.Phone,
		Status:    1,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
