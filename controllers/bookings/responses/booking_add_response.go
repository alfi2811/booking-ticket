package responses

import (
	"booking-ticket/business/bookings"
	"time"
)

type BookingAddResponse struct {
	Id             int       `json:"id"`
	UserId         int       `json:"userId"`
	TimeScheduleId int       `json:"timeScheduleId"`
	NumberSeat     string    `json:"numberSeat"`
	Quantity       int       `json:"quantity"`
	TotalPrice     int       `json:"totalPrice"`
	QrCode         string    `json:"qrcode"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func FromDomain(domain bookings.Domain) BookingAddResponse {
	return BookingAddResponse{
		Id:             domain.ID,
		UserId:         domain.UserId,
		TimeScheduleId: domain.TimeScheduleId,
		NumberSeat:     domain.NumberSeat,
		Quantity:       domain.Quantity,
		TotalPrice:     domain.TotalPrice,
		QrCode:         domain.QrCode,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}
