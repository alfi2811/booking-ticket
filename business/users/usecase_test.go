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
}

func TestRegister(t *testing.T) {
	setup()
	userRepository.On("Register",
		mock.Anything,
		mock.AnythingOfType("Domain")).Return(userDomain, nil).Once()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		admin, err := userService.Register(context.Background(), userDomain)

		assert.Nil(t, err)
		assert.NotEqual(t, "aaa@gmail.com", admin.Email)
	})
}

func TestGetUser(t *testing.T) {
	setup()
	userRepository.On("GetUser",
		mock.Anything).Return([]users.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		_, err := userService.GetUser(context.Background())

		assert.Nil(t, err)
	})
}
