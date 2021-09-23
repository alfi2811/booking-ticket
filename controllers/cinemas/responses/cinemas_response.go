package responses

import (
	"booking-ticket/business/cinemas"
	"time"
)

type CinemaResponse struct {
	ID        int    `json:"id" `
	Name      string `json:"name"`
	Location  string `json:"location"`
	Maps      string `json:"maps"`
	Phone     string `json:"phone"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDomain(domain cinemas.Domain) CinemaResponse {
	return CinemaResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		Location:  domain.Location,
		Maps:      domain.Maps,
		Phone:     domain.Phone,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
