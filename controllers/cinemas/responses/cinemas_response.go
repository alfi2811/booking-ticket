package responses

import (
	"booking-ticket/business/cinemas"
	"time"
)

type DetailResponse struct {
	Id       string    `json:"scheduleId"`
	IdMovie  string    `json:"movieId"`
	Title    string    `json:"title"`
	Poster   string    `json:"poster"`
	Duration string    `json:"duration"`
	Date     time.Time `json:"date"`
	Price    int       `json:"price"`
}

type CinemaResponse struct {
	ID        int    `json:"id" `
	Name      string `json:"name"`
	Location  string `json:"location"`
	Maps      string `json:"maps"`
	Phone     string `json:"phone"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CinemaDetaillResponse struct {
	CinemaResponse CinemaResponse   `json:"cinema"`
	CinemaDetail   []DetailResponse `json:"scehdule"`
}

func FromDomain(domain cinemas.Domain) CinemaResponse {
	return CinemaResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		Location:  domain.Location,
		Maps:      domain.Maps,
		Phone:     domain.Phone,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(data []cinemas.Domain) (result []CinemaResponse) {
	result = []CinemaResponse{}
	for _, cinema := range data {
		result = append(result, FromDomain(cinema))
	}
	return
}
func FromListDetail(data []cinemas.CinemaDetail) (result []DetailResponse) {
	result = []DetailResponse{}
	for _, detail := range data {
		result = append(result, FromDetail(detail))
	}
	return
}

func FromDetail(domain cinemas.CinemaDetail) DetailResponse {
	return DetailResponse{
		Id:       domain.Id,
		IdMovie:  domain.IdMovie,
		Title:    domain.Title,
		Poster:   domain.Poster,
		Duration: domain.Duration,
		Date:     domain.Date,
		Price:    domain.Price,
	}
}
func FromDomainDetail(domain cinemas.Domain) CinemaDetaillResponse {
	cinema := CinemaResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		Location:  domain.Location,
		Maps:      domain.Maps,
		Phone:     domain.Phone,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
	return CinemaDetaillResponse{
		CinemaResponse: cinema,
		CinemaDetail:   FromListDetail(domain.DetailCinema),
	}
}
