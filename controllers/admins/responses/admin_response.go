package responses

import (
	"booking-ticket/business/admins"
	"time"
)

type AdminResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain admins.Domain) AdminResponse {
	return AdminResponse{
		ID:        domain.ID,
		Fullname:  domain.Fullname,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(data []admins.Domain) (result []AdminResponse) {
	result = []AdminResponse{}
	for _, admin := range data {
		result = append(result, FromDomain(admin))
	}
	return
}
