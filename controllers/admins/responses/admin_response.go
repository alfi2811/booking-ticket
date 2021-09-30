package responses

import (
	"booking-ticket/business/admins"
	"time"
)

type AdminAddResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomainAdd(domain admins.Domain) AdminAddResponse {
	return AdminAddResponse{
		ID:        domain.ID,
		Fullname:  domain.Fullname,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
