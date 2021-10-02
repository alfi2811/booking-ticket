package admins_test

import (
	_middleware "booking-ticket/app/middlewares"
	"booking-ticket/business/admins"
	_mockRepository "booking-ticket/business/admins/mocks"
	"context"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var adminRepository _mockRepository.Repository
var adminService admins.Usecase
var adminDomain admins.Domain

func setup() {
	configJWTT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}
	adminService = admins.NewAdminUsecase(&adminRepository, time.Hour*1, configJWTT)
	adminDomain = admins.Domain{
		ID:       1,
		Email:    "al@gmail.com",
		Password: "123",
		Fullname: "alfi",
		Status:   1,
		Token:    "12",
	}
}

func TestRegister(t *testing.T) {
	setup()
	adminRepository.On("Register",
		mock.Anything,
		mock.AnythingOfType("Domain")).Return(adminDomain, nil).Once()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		admin, err := adminService.Register(context.Background(), adminDomain)

		assert.Nil(t, err)
		assert.NotEqual(t, "aaa@gmail.com", admin.Email)
	})
}

func TestLogin(t *testing.T) {
	setup()
	adminRepository.On("Login",
		mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).Return(adminDomain, nil).Once()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		admin, err := adminService.Login(context.Background(), "al@gmail.com", "123")

		assert.Nil(t, err)
		assert.Equal(t, "alfi", admin.Fullname)
	})
}
