package movies

import "time"

type MovieCreate struct {
	Title        string    `json:"title" form:"title" gorm:"unique"`
	Poster       string    `json:"poster" form:"poster"`
	ReleaseDate  time.Time `json:"releaseDate" form:"releaseDate"`
	Duration     string    `json:"duration" form:"duration"`
	MovieTrailer string    `json:"movieTrailer" form:"movieTrailer"`
	Players      string    `json:"players" form:"players"`
}
