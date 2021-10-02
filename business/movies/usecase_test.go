package movies_test

import (
	"booking-ticket/business/movies"
	_mockRepository "booking-ticket/business/movies/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var movieRepository _mockRepository.Repository
var movieService movies.Usecase
var movieDomain movies.Domain

func setup() {
	movieService = movies.NewMovieUsecase(&movieRepository, time.Hour*1)
	movieDomain = movies.Domain{
		ID:           1,
		Title:        "Fast and Furious",
		Poster:       "cdn.hdfhdjdjd",
		Duration:     "2 Jam 2 Minutes",
		MovieTrailer: "youtube.com/wjjwjw",
		Status:       1,
	}
}

func TestListMovie(t *testing.T) {
	setup()
	movieRepository.On("ListMovie",
		mock.Anything).Return([]movies.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid List Cinema", func(t *testing.T) {
		_, err := movieService.ListMovie(context.Background())

		assert.Nil(t, err)
	})
}

func TestCinemaDetail(t *testing.T) {
	setup()
	movieRepository.On("DetailMovie",
		mock.Anything,
		mock.AnythingOfType("int")).Return(movies.Domain{}, nil).Once()

	t.Run("Test Case 1 | Valid Detail", func(t *testing.T) {
		_, err := movieService.DetailMovie(context.Background(), 1)

		assert.Nil(t, err)
	})
}
