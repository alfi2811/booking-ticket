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

func ToListDomain(data []Movies) (result []movies.Domain) {
	result = []movies.Domain{}
	for _, movie := range data {
		result = append(result, movie.ToDomain())
	}
	return
}

func FromDomain(domain movies.Domain) Movies {
	return Movies{
		Title:        domain.Title,
		Poster:       domain.Poster,
		ReleaseDate:  domain.ReleaseDate,
		Duration:     domain.Duration,
		MovieTrailer: domain.MovieTrailer,
		Players:      domain.Players,
		Status:       1,
	}
}
