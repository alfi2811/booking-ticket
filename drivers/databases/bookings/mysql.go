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

func (rep *MysqlBookingRepository) ListBooking(ctx context.Context) ([]bookings.Domain, error) {
	var listBooking []Bookings
	result := rep.Conn.Find(&listBooking)

	if result.Error != nil {
		return []bookings.Domain{}, result.Error
	}

	return ToListDomain(listBooking), nil
}

func (rep *MysqlBookingRepository) ListBookingUser(ctx context.Context, userId int) ([]bookings.DomainJoin, error) {
	var listBooking []Bookings
	var listBookingUser []bookings.DomainJoin

	result := rep.Conn.Model(&listBooking).Select("bookings.id, movies.title as title_movie, movies.poster as poster_movie, cinemas.name as name_cinema, time_schedules.time_movie as time").Joins("JOIN time_schedules on bookings.time_schedule_id = time_schedules.id").Joins("JOIN schedules on time_schedules.schedule_id = schedules.id").Joins("JOIN movies on schedules.movie_id = movies.id").Joins("JOIN cinemas on schedules.cinema_id = cinemas.id").Where("bookings.user_id = ?", userId).Scan(&listBookingUser)

	if result.Error != nil {
		return []bookings.DomainJoin{}, result.Error
	}

	return listBookingUser, nil
}
