package cinemas

import (
	"booking-ticket/business/cinemas"
	"booking-ticket/drivers/databases/schedules"
	"time"
)

type Cinemas struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	Location  string
	Maps      string
	Phone     string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Schedule  []schedules.Schedules `gorm:"foreignKey:CinemaId;references:ID"`
}

func (cinema *Cinemas) ToDomain() cinemas.Domain {
	return cinemas.Domain{
		ID:        cinema.ID,
		Name:      cinema.Name,
		Location:  cinema.Location,
		Maps:      cinema.Maps,
		Phone:     cinema.Phone,
		Status:    cinema.Status,
		CreatedAt: cinema.CreatedAt,
		UpdatedAt: cinema.UpdatedAt,
	}
}

func ToListDomain(data []Cinemas) (result []cinemas.Domain) {
	result = []cinemas.Domain{}
	for _, cinema := range data {
		result = append(result, cinema.ToDomain())
	}
	return
}

func FromDomain(domain cinemas.Domain) Cinemas {
	return Cinemas{
		ID:       domain.ID,
		Name:     domain.Name,
		Location: domain.Location,
		Maps:     domain.Maps,
		Phone:    domain.Phone,
		Status:   1,
	}
}
