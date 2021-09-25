package bookings

import (
	"booking-ticket/business/bookings"
	"context"

	"gorm.io/gorm"
)

type MysqlBookingRepository struct {
	Conn *gorm.DB
}

func NewMysqlBookingsRepository(conn *gorm.DB) bookings.Repository {
	return &MysqlBookingRepository{
		Conn: conn,
	}
}

func (rep *MysqlBookingRepository) AddBooking(ctx context.Context, domain bookings.Domain) (bookings.Domain, error) {
	bookingDB := FromDomain(domain)

	result := rep.Conn.Create(&bookingDB)

	if result.Error != nil {
		return bookings.Domain{}, result.Error
	}

	return bookingDB.ToDomain(), nil
}
