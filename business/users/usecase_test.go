package users_test

import (
	_middleware "booking-ticket/app/middlewares"
	"booking-ticket/business/users"
	_mockRepository "booking-ticket/business/users/mocks"
	"context"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _mockRepository.Repository
var userService users.Usecase
var userDomain users.Domain

func setup() {
	configJWTT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}
	userService = users.NewUserUsecase(&userRepository, time.Hour*1, configJWTT)
	userDomain = users.Domain{
		ID:       1,
		Email:    "al@gmail.com",
		Password: "123",
		Fullname: "alfi",
		Gender:   "Laki-Laki",
		Phone:    "08129217272",
		Status:   1,
		Token:    "12",
	}
}

func TestLogin(t *testing.T) {
	setup()
	userRepository.On("Login",
		mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).Return(userDomain, nil).Once()

	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		admin, err := userService.Login(context.Background(), "al@gmail.com", "123")

		assert.Nil(t, err)
		assert.Equal(t, "alfi", admin.Fullname)
	})
	t.Run("Test Case 2 | Invalid Email Empty", func(t *testing.T) {
		_, err := userService.Login(context.Background(), "", "123")

		assert.NotNil(t, err)
	})
	t.Run("Test Case 3 | Invalid Password Empty", func(t *testing.T) {
		_, err := userService.Login(context.Background(), "al@gmail.com", "")

		assert.NotNil(t, err)
	})
}

func TestRegister(t *testing.T) {
	setup()
	userRepository.On("Register",
		mock.Anything,
		mock.AnythingOfType("Domain")).Return(userDomain, nil).Once()
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		admin, err := userService.Register(context.Background(), userDomain)

		assert.Nil(t, err)
		assert.NotEqual(t, "aaa@gmail.com", admin.Email)
	})
	t.Run("Test Case 2 | Invalid Register Field Empty", func(t *testing.T) {
		_, err := userService.Register(context.Background(), users.Domain{
			Email:    "al@gmail.com",
			Password: "",
			Fullname: "alfi",
			Gender:   "Laki-Laki",
			Phone:    "0891828282",
		})

		assert.NotNil(t, err)
	})
}

func TestGetUser(t *testing.T) {
	setup()
	userRepository.On("GetUser",
		mock.Anything).Return([]users.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid Get User", func(t *testing.T) {
		_, err := userService.GetUser(context.Background())

		assert.Nil(t, err)
	})
}

func TestListUserBooking(t *testing.T) {
	setup()
	userRepository.On("ListUserBooking",
		mock.Anything).Return([]users.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid List User Booking", func(t *testing.T) {
		_, err := userService.ListUserBooking(context.Background())

		assert.Nil(t, err)
	})
}
