package schedules

import (
	"booking-ticket/business/schedules"
	"context"

	"gorm.io/gorm"
)

type MysqlSchedulesRepository struct {
	Conn *gorm.DB
}

func NewMysqlScheduleRepository(conn *gorm.DB) schedules.Repository {
	return &MysqlSchedulesRepository{
		Conn: conn,
	}
}

func (rep *MysqlSchedulesRepository) AddSchedule(ctx context.Context, domain schedules.Domain) (schedules.Domain, error) {
	cinemaDB := FromDomain(domain)

	result := rep.Conn.Create(&cinemaDB)

	if result.Error != nil {
		return schedules.Domain{}, result.Error
	}

	return cinemaDB.ToDomain(), nil
}
