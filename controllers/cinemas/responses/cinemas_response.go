package responses

import (
	"booking-ticket/business/cinemas"
	"time"
)

type CinemaDetailReponse struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Poster   string    `json:"poster"`
	Duration string    `json:"duration"`
	Date     time.Time `json:"date"`
	Price    int       `json:"price"`
}

type CinemaResponse struct {
	ID          int                    `json:"id" `
	Name        string                 `json:"name"`
	Location    string                 `json:"location"`
	Maps        string                 `json:"maps"`
	Phone       string                 `json:"phone"`
	MovieDetail []cinemas.CinemaDetail `json:"movie"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// type CinemaDetaillResponose struct {
// 	CinemaResponse,
// 	m
// }

func FromDomain(domain cinemas.Domain) CinemaResponse {
	return CinemaResponse{
		ID:          domain.ID,
		Name:        domain.Name,
		Location:    domain.Location,
		Maps:        domain.Maps,
		Phone:       domain.Phone,
		MovieDetail: domain.DetailCinema,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
