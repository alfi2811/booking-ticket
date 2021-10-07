package responses

import (
	"booking-ticket/business/users"
	"booking-ticket/controllers/bookings/responses"
	"booking-ticket/drivers/databases/bookings"
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
}

type UserBookingResponse struct {
	User     UserResponse                `json:"user"`
	Bookings []responses.BookingResponse `json:"bookings"`
}

type UserLoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		ID:        domain.ID,
		Fullname:  domain.Fullname,
		Email:     domain.Email,
		Gender:    domain.Gender,
		Phone:     domain.Phone,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainUserBook(domain users.Domain) UserBookingResponse {
	user := FromDomain(domain)
	bookingDB := bookings.ToListDomain(domain.Booking)
	booking := responses.FromListDomain(bookingDB)
	return UserBookingResponse{
		User:     user,
		Bookings: booking,
	}
}

func FromListUserBookDomain(data []users.Domain) (result []UserBookingResponse) {
	result = []UserBookingResponse{}
	for _, user := range data {
		result = append(result, FromDomainUserBook(user))
	}
	return
}

func FromListDomain(data []users.Domain) (result []UserResponse) {
	result = []UserResponse{}
	for _, user := range data {
		result = append(result, FromDomain(user))
	}
	return
}

func FromDomainLogin(domain users.Domain) UserLoginResponse {
	user := UserResponse{
		ID:        domain.ID,
		Fullname:  domain.Fullname,
		Email:     domain.Email,
		Gender:    domain.Gender,
		Phone:     domain.Phone,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
	return UserLoginResponse{
		User:  user,
		Token: domain.Token,
	}
}
