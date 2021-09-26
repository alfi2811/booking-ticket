package requests

import (
	"booking-ticket/business/bookings"
)

type BookingAdd struct {
	UserId         int    `json:"userId"`
	TimeScheduleId int    `json:"timeScheduleId"`
	NumberSeat     string `json:"numberSeat"`
	Quantity       int    `json:"quantity"`
	TotalPrice     int    `json:"totalPrice"`
}

func (booking *BookingAdd) ToDomain() bookings.Domain {
	return bookings.Domain{
		UserId:         booking.UserId,
		TimeScheduleId: booking.TimeScheduleId,
		NumberSeat:     booking.NumberSeat,
		Quantity:       booking.Quantity,
		TotalPrice:     booking.TotalPrice,
	}
}
