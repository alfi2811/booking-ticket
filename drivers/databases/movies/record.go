package movies

import (
	"booking-ticket/business/movies"
	"booking-ticket/drivers/databases/schedules"
	"time"
)

type Movies struct {
	ID           int    `gorm:"primaryKey"`
	Title        string `gorm:"unique"`
	Poster       string
	ReleaseDate  time.Time
	Duration     string
	MovieTrailer string
	Players      string
	Status       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Schedule     []schedules.Schedules `gorm:"foreignKey:MovieId;references:ID"`
}
type UsersArr []Movies

func (movie *Movies) ToDomain() movies.Domain {
	return movies.Domain{
		ID:           movie.ID,
		Title:        movie.Title,
		Poster:       movie.Poster,
		ReleaseDate:  movie.ReleaseDate,
		Duration:     movie.Duration,
		MovieTrailer: movie.MovieTrailer,
		Players:      movie.Players,
		Status:       movie.Status,
		CreatedAt:    movie.CreatedAt,
		UpdatedAt:    movie.UpdatedAt,
	}
}

func FromDomain(domain movies.Domain) Movies {
	return Movies{
		ID:           domain.ID,
		Title:        domain.Title,
		Poster:       domain.Poster,
		ReleaseDate:  domain.ReleaseDate,
		Duration:     domain.Duration,
		MovieTrailer: domain.MovieTrailer,
		Players:      domain.Players,
		Status:       domain.Status,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
