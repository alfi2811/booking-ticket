package responses

import (
	"booking-ticket/business/admins"
)

type AdminResponseLogin struct {
	Admin AdminResponse `json:"admin"`
	Token string        `json:"token"`
}

func FromDomainLogin(domain admins.Domain) AdminResponseLogin {
	admin := AdminResponse{
		ID:        domain.ID,
		Fullname:  domain.Fullname,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
	return AdminResponseLogin{
		Admin: admin,
		Token: domain.Token,
	}
}
