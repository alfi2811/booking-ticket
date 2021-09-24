package schedules

import (
	"booking-ticket/business/schedules"
	"context"
	"fmt"

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
	// result := rep.Conn.Find(&ListSchedule)
	result := rep.Conn.Model(&ListSchedule).Select("movies.title ,schedules.date, schedules.price").Joins("left join movies on schedules.movie_id = movies.id")

	fmt.Println(result)

	if result.Error != nil {
		return []schedules.Domain{}, result.Error
	}

	return ToListDomain(ListSchedule), nil
}
