package cinemas

import (
	"booking-ticket/business/cinemas"
	"booking-ticket/drivers/databases/schedules"
	"context"

	"gorm.io/gorm"
)

type MysqlCinemaRepository struct {
	Conn *gorm.DB
}

func NewMysqlCinemaRepository(conn *gorm.DB) cinemas.Repository {
	return &MysqlCinemaRepository{
		Conn: conn,
	}
}

func (rep *MysqlCinemaRepository) AddCinema(ctx context.Context, domain cinemas.Domain) (cinemas.Domain, error) {
	cinemaDB := FromDomain(domain)

	result := rep.Conn.Create(&cinemaDB)

	if result.Error != nil {
		return cinemas.Domain{}, result.Error
	}

	return cinemaDB.ToDomain(), nil
}

func (rep *MysqlCinemaRepository) ListCinema(ctx context.Context) ([]cinemas.Domain, error) {
	var listCinema []Cinemas
	result := rep.Conn.Find(&listCinema)

	if result.Error != nil {
		return []cinemas.Domain{}, result.Error
	}

	return ToListDomain(listCinema), nil
}

func (rep *MysqlCinemaRepository) CinemaDetail(ctx context.Context, cinemaId int) (cinemas.Domain, error) {
	var cinema Cinemas
	var schedules []schedules.Schedules
	var detail []cinemas.CinemaDetail
	// var domainUser []users.Domain
	result := rep.Conn.Where("id = ?", cinemaId).Preload("Schedule").First(&cinema)
	resultt := rep.Conn.Model(&schedules).Select("schedules.id as id_schedule, movies.id as id_movie , movies.title, movies.poster, movies.duration, schedules.date, schedules.price").Joins("JOIN movies on schedules.movie_id = movies.id").Joins("JOIN cinemas on schedules.cinema_id = cinemas.id").Where("cinemas.id = ?", cinemaId).Scan(&detail)

	if result.Error != nil && resultt.Error != nil {
		return cinemas.Domain{}, result.Error
	}
	cinemaDomain := cinema.ToDomain()
	cinemaDomain.DetailCinema = detail
	return cinemaDomain, nil
}
