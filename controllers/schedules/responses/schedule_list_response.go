package responses

import (
	"time"
)

type ScheduleListResponse struct {
	ID    int       `json:"id" `
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
	Price int       `json:"price"`
}
