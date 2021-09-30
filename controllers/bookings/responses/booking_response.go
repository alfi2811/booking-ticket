package responses

import (
	"booking-ticket/business/bookings"
	"time"
)

type BookingResponse struct {
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

func FromDomain(domain bookings.Domain) BookingResponse {
	return BookingResponse{
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

func FromListDomain(data []bookings.Domain) (result []BookingResponse) {
	result = []BookingResponse{}
	for _, movie := range data {
		result = append(result, FromDomain(movie))
	}
	return
}
