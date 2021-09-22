package schedules

import "time"

type ScheduleCreate struct {
	MovieId  int       `json:"movieId" form:"movieId"`
	CinemaId int       `json:"cinemaId" form:"cinemaId"`
	Date     time.Time `json:"date" form:"date"`
	Price    int       `json:"price" form:"price"`
}
