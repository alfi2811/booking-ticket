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
	var movieDB Movies

	movieDB.Title = domain.Title
	movieDB.Poster = domain.Poster
	movieDB.ReleaseDate = domain.ReleaseDate
	movieDB.Duration = domain.Duration
	movieDB.MovieTrailer = domain.MovieTrailer
	movieDB.Players = domain.Players
	movieDB.Status = 1

	result := rep.Conn.Create(&movieDB)

	if result.Error != nil {
		return movies.Domain{}, result.Error
	}

	return movieDB.ToDomain(), nil
}
