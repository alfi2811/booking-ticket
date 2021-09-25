package requests

import "booking-ticket/business/users"

type UserRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
}

func (user *UserRegister) ToDomain() users.Domain {
	return users.Domain{
		Email:    user.Email,
		Password: user.Password,
		Fullname: user.Fullname,
		Gender:   user.Gender,
		Phone:    user.Phone,
	}
}
