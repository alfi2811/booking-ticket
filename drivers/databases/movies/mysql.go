package movies

import (
	"booking-ticket/business/movies"
	"context"

	"gorm.io/gorm"
)

type MysqlMovieRepository struct {
	Conn *gorm.DB
}

func NewMysqlMovieRepository(conn *gorm.DB) movies.Repository {
	return &MysqlMovieRepository{
		Conn: conn,
	}
}

func (rep *MysqlMovieRepository) AddMovie(ctx context.Context, domain movies.Domain) (movies.Domain, error) {
	movieDB := FromDomain(domain)

	result := rep.Conn.Create(&movieDB)

	if result.Error != nil {
		return movies.Domain{}, result.Error
	}

	return movieDB.ToDomain(), nil
}

func (rep *MysqlMovieRepository) ListMovie(ctx context.Context) ([]movies.Domain, error) {
	var listMovies []Movies
	// var domainUser []users.Domain
	result := rep.Conn.Find(&listMovies)

	if result.Error != nil {
		return []movies.Domain{}, result.Error
	}

	return ToListDomain(listMovies), nil
}

func (rep *MysqlMovieRepository) DetailMovie(ctx context.Context, idMovie int) (movies.Domain, error) {
	var movie Movies

	result := rep.Conn.Where("id = ?", idMovie).First(&movie)
	if result.Error != nil {
		return movies.Domain{}, result.Error
	}

	return movie.ToDomain(), nil
}
