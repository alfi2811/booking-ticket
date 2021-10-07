package schedules

import (
	"booking-ticket/business/schedules"
	"booking-ticket/drivers/databases/timeSchedules"
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

func (rep *MysqlSchedulesRepository) ListSchedule(ctx context.Context) ([]schedules.Domain, error) {
	var ListSchedule []Schedules
	var schedulesData []ListSchedules
	result := rep.Conn.Model(&ListSchedule).Select("schedules.id, cinemas.name, cinemas.maps, cinemas.phone, movies.title, movies.poster, movies.duration, schedules.date, schedules.price").Joins("left join movies on schedules.movie_id = movies.id").Joins("JOIN cinemas on schedules.cinema_id = cinemas.id").Scan(&schedulesData)

	if result.Error != nil {
		return []schedules.Domain{}, result.Error
	}

	return ToListDomainJoin(schedulesData), nil
}

func (rep *MysqlSchedulesRepository) DetailTimeSchedule(ctx context.Context, idSchedule int) (schedules.Domain, error) {
	var ListSchedule []Schedules
	var scheduleData ListSchedules
	var timeDB []timeSchedules.TimeSchedules
	resultTime := rep.Conn.Where("schedule_id = ?", idSchedule).Find(&timeDB)

	result := rep.Conn.Model(&ListSchedule).Select("schedules.id, cinemas.name, cinemas.maps, cinemas.phone, movies.title, movies.poster, movies.duration, schedules.date, schedules.price").Joins("left join movies on schedules.movie_id = movies.id").Joins("JOIN cinemas on schedules.cinema_id = cinemas.id").Where("schedules.id = ?", idSchedule).Find(&scheduleData)

	if result.Error != nil && resultTime.Error != nil {
		return schedules.Domain{}, result.Error
	}
	scheduleDomain := scheduleData.ToDomainJoin()
	scheduleDomain.Time = timeSchedules.ToListDomain(timeDB)
	return scheduleDomain, nil
}
