package admins_test

import (
	_middleware "booking-ticket/app/middlewares"
	"booking-ticket/business"
	"booking-ticket/business/admins"
	_mockRepository "booking-ticket/business/admins/mocks"
	"booking-ticket/helpers/encrypt"
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

	t.Run("Test Case 2 | invalid empty field", func(t *testing.T) {
		_, err := adminService.Register(context.Background(), admins.Domain{
			Email:    "al@gmail.com",
			Password: "",
			Fullname: "alfi",
		})

		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	setup()
	adminDomain.Password, _ = encrypt.Hash(adminDomain.Password)
	adminRepository.On("Login",
		mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).Return(adminDomain, nil).Once()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		admin, err := adminService.Login(context.Background(), "al@gmail.com", "123")

		assert.Nil(t, err)
		assert.Equal(t, adminDomain.Fullname, admin.Fullname)
		assert.Equal(t, adminDomain.Password, admin.Password)
	})

	t.Run("Test Case 2 | Invalid Empty Email", func(t *testing.T) {
		_, err := adminService.Login(context.Background(), "", "123")

		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Invalid Empty Password", func(t *testing.T) {
		_, err := adminService.Login(context.Background(), "al@gmail.com", "")

		assert.NotNil(t, err)
	})

	t.Run("Test Case 4 | Invalid Wrong Password", func(t *testing.T) {
		adminRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(adminDomain, business.ErrPasswordWrong).Once()
		_, err := adminService.Login(context.Background(), "all@gmail.com", "12345")

		assert.NotNil(t, err)
	})
}

func TestGetAdmin(t *testing.T) {
	setup()
	adminRepository.On("GetAdmin",
		mock.Anything).Return([]admins.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid Get Admin", func(t *testing.T) {
		_, err := adminService.GetAdmin(context.Background())

		assert.Nil(t, err)
	})
}
