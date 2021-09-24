package users

import (
	"booking-ticket/business/admins"
	"time"
)

type Admins struct {
	ID        int    `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Password  string
	Fullname  string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *Admins) ToDomain() admins.Domain {
	return admins.Domain{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		Fullname:  user.Fullname,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
func ToListDomain(data []Admins) (result []admins.Domain) {
	result = []admins.Domain{}
	for _, admin := range data {
		result = append(result, admin.ToDomain())
	}
	return
}

func FromDomain(domain admins.Domain) Admins {
	return Admins{
		ID:        domain.ID,
		Email:     domain.Email,
		Password:  domain.Password,
		Fullname:  domain.Fullname,
		Status:    1,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
