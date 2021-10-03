package schedules_test

import (
	"booking-ticket/business/schedules"
	_mockRepository "booking-ticket/business/schedules/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var scheduleRepository _mockRepository.Repository
var scheduleService schedules.Usecase
var scheduleDomain schedules.Domain

func setup() {
	scheduleService = schedules.NewScheduleUsecase(&scheduleRepository, time.Hour*1)
	scheduleDomain = schedules.Domain{
		ID:       1,
		CinemaId: 1,
		MovieId:  1,
		Date:     time.Now(),
		Price:    40000,
		Status:   1,
	}
}

func TestAddSchedule(t *testing.T) {
	setup()
	scheduleRepository.On("AddSchedule",
		mock.Anything,
		mock.AnythingOfType("Domain")).Return(schedules.Domain{}, nil).Once()

	t.Run("Test Case 1 | Valid Add Schedules", func(t *testing.T) {
		_, err := scheduleService.AddSchedule(context.Background(), scheduleDomain)

		assert.Nil(t, err)
	})
	t.Run("Test Case 2 | Invalid Add Schedules Relation ID Empty", func(t *testing.T) {
		_, err := scheduleService.AddSchedule(context.Background(), schedules.Domain{
			CinemaId: 1,
			MovieId:  0,
			Date:     time.Now(),
			Price:    40000,
		})

		assert.NotNil(t, err)
	})
	t.Run("Test Case 3 | Invalid Add Schedules Field Empty", func(t *testing.T) {
		_, err := scheduleService.AddSchedule(context.Background(), schedules.Domain{
			CinemaId: 1,
			MovieId:  1,
			Date:     time.Now(),
			Price:    0,
		})

		assert.NotNil(t, err)
	})
}

func TestListSchedule(t *testing.T) {
	setup()
	scheduleRepository.On("ListSchedule",
		mock.Anything).Return([]schedules.Domain{}, nil).Once()

	t.Run("Test Case 1 | Valid List Schedules", func(t *testing.T) {
		_, err := scheduleService.ListSchedule(context.Background())

		assert.Nil(t, err)
	})
}

func TestDetailTimeSchedule(t *testing.T) {
	setup()
	scheduleRepository.On("DetailTimeSchedule",
		mock.Anything,
		mock.AnythingOfType("int")).Return(schedules.Domain{}, nil).Once()

	t.Run("Test Case 1 | Valid Detail", func(t *testing.T) {
		_, err := scheduleService.DetailTimeSchedule(context.Background(), scheduleDomain.ID)

		assert.Nil(t, err)
	})
}
