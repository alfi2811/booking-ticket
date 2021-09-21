package schedules

import "time"

type ScheduleResponse struct {
	Title string    `json:"title" form:"title"`
	Date  time.Time `json:"date" form:"date"`
	Price int       `json:"price" form:"price"`
}
