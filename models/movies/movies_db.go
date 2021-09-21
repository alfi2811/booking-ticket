package movies

import "time"

type Movie struct {
	ID           int       `json:"id" form:"id" gorm:"primaryKey" `
	Title        string    `json:"title" form:"title" gorm:"unique"`
	Poster       string    `json:"poster" form:"poster"`
	ReleaseDate  time.Time `json:"releaseDate" form:"releaseDate"`
	Duration     string    `json:"duration" form:"duration"`
	MovieTrailer string    `json:"movieTrailer" form:"movieTrailer"`
	Players      string    `json:"players" form:"players"`
	Status       int       `json:"status" form:"status"`
	CreatedAt    time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt" form:"updatedAt"`
}
