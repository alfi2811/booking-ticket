package cinemas

import (
	"booking-ticket/business/cinemas"
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
