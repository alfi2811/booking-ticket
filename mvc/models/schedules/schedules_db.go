package schedules

import "time"

type Schedule struct {
	ID        int       `json:"id" form:"id" gorm:"primaryKey" `
	MovieId   int       `json:"movie_id" form:"movie_id"`
	CinemaId  int       `json:"cinema_id" form:"cinema_id"`
	Date      time.Time `json:"date" form:"date"`
	Price     int       `json:"price" form:"price"`
	Status    int       `json:"status" form:"status"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
}
