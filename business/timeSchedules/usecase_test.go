package timeSchedules_test

import (
	"booking-ticket/business/timeSchedules"
	_mockRepository "booking-ticket/business/timeSchedules/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var timeScheduleRepository _mockRepository.Repository
var timeScheduleService timeSchedules.Usecase
var timeScheduleDomain timeSchedules.Domain

func setup() {
	timeScheduleService = timeSchedules.NewTimeScheduleUsecase(&timeScheduleRepository, time.Hour*1)
	timeScheduleDomain = timeSchedules.Domain{
		ID:         1,
		ScheduleId: 1,
		TimeMovie:  time.Now(),
	}
}

func TestAddScheduleTime(t *testing.T) {
	setup()
	timeScheduleRepository.On("AddScheduleTime",
		mock.Anything,
		mock.AnythingOfType("Domain")).Return(timeSchedules.Domain{}, nil).Once()

	t.Run("Test Case 1 | Valid Add Schedules", func(t *testing.T) {
		_, err := timeScheduleService.AddScheduleTime(context.Background(), timeScheduleDomain)

		assert.Nil(t, err)
	})
	t.Run("Test Case 2 | Invalid Add Schedules field empty", func(t *testing.T) {
		_, err := timeScheduleService.AddScheduleTime(context.Background(), timeSchedules.Domain{
			ScheduleId: 0,
			TimeMovie:  time.Now(),
		})

		assert.NotNil(t, err)
	})
}

func TestDetailScheduleTime(t *testing.T) {
	setup()
	timeScheduleRepository.On("DetailScheduleTime",
		mock.Anything,
		mock.AnythingOfType("int")).Return([]timeSchedules.Domain{}, nil).Once()

	t.Run("Test Case 1 | Valid Detail", func(t *testing.T) {
		_, err := timeScheduleService.DetailScheduleTime(context.Background(), timeScheduleDomain.ID)

		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid Detail No params", func(t *testing.T) {
		_, err := timeScheduleService.DetailScheduleTime(context.Background(), 0)

		assert.NotNil(t, err)
	})
}
