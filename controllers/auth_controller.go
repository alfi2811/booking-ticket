package controllers

import (
	"booking-ticket/config"
	"booking-ticket/models/auth"
	"booking-ticket/models/response"
	"booking-ticket/models/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	var userDB users.User
	userLogin := auth.AuthLoginUser{}
	c.Bind(&userLogin)
	// login
	result := config.DB.Where("email = ? AND password = ?", userLogin.Email, userLogin.Password).Find(&userDB)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error, Please try again later!",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Login Successfull",
		Data:    userDB,
	})
}

func RegisterController(c echo.Context) error {
	var userRegister auth.AuthRegisterUser
	c.Bind(&userRegister)

	// validasi
	if userRegister.Fullname == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Fullname still empty",
			Data:    nil,
		})
	}

	var userDB users.User
	userDB.Email = userRegister.Email
	userDB.Password = userRegister.Password
	userDB.Fullname = userRegister.Fullname
	userDB.Gender = userRegister.Gender
	userDB.Phone = userRegister.Phone
	userDB.Status = 1

	result := config.DB.Create(&userDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error, Please try again later!",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Registeration Successfully",
		Data:    userDB,
	})
}
