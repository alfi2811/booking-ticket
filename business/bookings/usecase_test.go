package bookings_test

import (
	"booking-ticket/business/bookings"
	_mockRepository "booking-ticket/business/bookings/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var bookingRepository _mockRepository.Repository
var bookingService bookings.Usecase
var bookingDomain bookings.Domain

func setup() {
	bookingService = bookings.NewBookingUsecase(&bookingRepository, time.Hour*1)
	bookingDomain = bookings.Domain{
		ID:             1,
		UserId:         1,
		TimeScheduleId: 1,
		NumberSeat:     "A9",
		Quantity:       1,
		TotalPrice:     1,
		Status:         1,
	}
}

func TestAddBooking(t *testing.T) {
	setup()
	bookingRepository.On("AddBooking",
		mock.Anything,
		mock.AnythingOfType("Domain")).Return(bookings.Domain{}, nil).Once()

	t.Run("Test Case 1 | Valid Add Booking", func(t *testing.T) {
		_, err := bookingService.AddBooking(context.Background(), bookingDomain)

		assert.Nil(t, err)
	})
	t.Run("Test Case 2 | Invalid Add Booking Field Empty", func(t *testing.T) {
		_, err := bookingService.AddBooking(context.Background(), bookings.Domain{
			UserId:         1,
			TimeScheduleId: 1,
			NumberSeat:     "",
			Quantity:       1,
			TotalPrice:     1,
		})
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3 | Invalid Add Booking Relation ID Empty", func(t *testing.T) {
		_, err := bookingService.AddBooking(context.Background(), bookings.Domain{
			UserId:         1,
			TimeScheduleId: 0,
			NumberSeat:     "A9",
			Quantity:       1,
			TotalPrice:     1,
		})
		assert.NotNil(t, err)
	})
}

func TestListBooking(t *testing.T) {
	setup()
	bookingRepository.On("ListBooking",
		mock.Anything).Return([]bookings.Domain{}, nil).Once()

	t.Run("Test Case 1 | Valid List Booking", func(t *testing.T) {
		_, err := bookingService.ListBooking(context.Background())

		assert.Nil(t, err)
	})
}

func TestListBookingUser(t *testing.T) {
	setup()
	bookingRepository.On("ListBookingUser",
		mock.Anything,
		mock.AnythingOfType("int")).Return([]bookings.DomainJoin{}, nil).Once()

	t.Run("Test Case 1 | Valid List Booking User", func(t *testing.T) {
		_, err := bookingService.ListBookingUser(context.Background(), 1)

		assert.Nil(t, err)
	})
}

func TestDetailTimeSchedule(t *testing.T) {
	setup()
	bookingRepository.On("DetailBooking",
		mock.Anything,
		mock.AnythingOfType("int")).Return(bookings.Domain{}, nil).Once()

	t.Run("Test Case 1 | Valid Detail", func(t *testing.T) {
		_, err := bookingService.DetailBooking(context.Background(), bookingDomain.ID)

		assert.Nil(t, err)
	})
}
