package requests

import "booking-ticket/business/admins"

type AdminRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

func (admin *AdminRegister) ToDomain() admins.Domain {
	return admins.Domain{
		Email:    admin.Email,
		Password: admin.Password,
		Fullname: admin.Fullname,
	}
}
