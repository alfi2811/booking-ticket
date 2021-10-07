package timeSchedules

import (
	"booking-ticket/business/timeSchedules"
	"context"

	"gorm.io/gorm"
)

type MysqlTimeSchedulesRepository struct {
	Conn *gorm.DB
}

func NewMysqlTimeScheduleRepository(conn *gorm.DB) timeSchedules.Repository {
	return &MysqlTimeSchedulesRepository{
		Conn: conn,
	}
}

func (rep *MysqlTimeSchedulesRepository) AddScheduleTime(ctx context.Context, domain timeSchedules.Domain) (timeSchedules.Domain, error) {
	timeDB := FromDomain(domain)

	result := rep.Conn.Create(&timeDB)

	if result.Error != nil {
		return timeSchedules.Domain{}, result.Error
	}

	return timeDB.ToDomain(), nil
}

func (rep *MysqlTimeSchedulesRepository) DetailScheduleTime(ctx context.Context, scheduleId int) ([]timeSchedules.Domain, error) {
	var timeDB []TimeSchedules
	result := rep.Conn.Where("schedule_id = ?", scheduleId).Find(&timeDB)

	if result.Error != nil {
		return []timeSchedules.Domain{}, result.Error
	}

	return ToListDomain(timeDB), nil
}
