package responses

import (
	"booking-ticket/business/users"
	"time"
)

type UserResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Token     string    `json:"token"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		ID:        domain.ID,
		Fullname:  domain.Fullname,
		Email:     domain.Email,
		Gender:    domain.Gender,
		Phone:     domain.Phone,
		Token:     domain.Token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
