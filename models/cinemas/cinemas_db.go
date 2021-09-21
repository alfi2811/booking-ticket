package cinemas

import (
	"booking-ticket/models/schedules"
	"time"
)

type Cinema struct {
	ID        int                  `json:"id" form:"id" gorm:"primaryKey" `
	Name      string               `json:"name" form:"name" gorm:"unique"`
	Location  string               `json:"location" form:"location"`
	Maps      string               `json:"maps" form:"maps"`
	Phone     string               `json:"phone" form:"phone"`
	Status    int                  `json:"status" form:"status"`
	CreatedAt time.Time            `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt" form:"updatedAt"`
	Schedule  []schedules.Schedule `gorm:"foreignKey:CinemaId;references:ID"`
}
