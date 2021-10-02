package cinemas_test

import (
	"booking-ticket/business/cinemas"
	_mockRepository "booking-ticket/business/cinemas/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var cinemaRepository _mockRepository.Repository
var cinemaService cinemas.Usecase
var userDomain cinemas.Domain

func setup() {
	cinemaService = cinemas.NewCinemaUsecase(&cinemaRepository, time.Hour*1)
	userDomain = cinemas.Domain{
		ID:       1,
		Name:     "Senayan XXI",
		Location: "Jl senayan no. 21",
		Maps:     "maps.google.com/1211",
		Phone:    "08129217272",
		Status:   1,
	}
}

func TestListCinema(t *testing.T) {
	setup()
	cinemaRepository.On("ListCinema",
		mock.Anything).Return([]cinemas.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		_, err := cinemaService.ListCinema(context.Background())

		assert.Nil(t, err)
	})
}

func TestCinemaDetail(t *testing.T) {
	setup()
	cinemaRepository.On("CinemaDetail",
		mock.Anything,
		mock.AnythingOfType("int")).Return(cinemas.Domain{}, nil).Once()

	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		_, err := cinemaService.CinemaDetail(context.Background(), 1)

		assert.Nil(t, err)
	})
}
