package controllers

import (
	"booking-ticket/config"
	"booking-ticket/middleware"
	"booking-ticket/models/auth"
	"booking-ticket/models/response"
	"booking-ticket/models/users"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LoginController(c echo.Context) error {
	var userDB users.User
	userLogin := auth.AuthLoginUser{}
	c.Bind(&userLogin)
	// login
	result := config.DB.Where("email = ? AND password = ?", userLogin.Email, userLogin.Password).First(&userDB)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusUnauthorized, response.BaseResponse{
				Code:    http.StatusUnauthorized,
				Message: "Email or Password Wrong",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error, Please try again later!",
			Data:    nil,
		})
	}

	token, err := middleware.CreateToken(userDB.ID, userDB.Fullname)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error, Please try again later!",
			Data:    nil,
		})
	}

	userResponse := users.UserResponse{userDB.ID, userDB.Email, userDB.Fullname, token}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Login Successfull",
		Data:    userResponse,
	})
}

func RegisterController(c echo.Context) error {
	var userRegister auth.AuthRegisterUser
	c.Bind(&userRegister)

	// validasi
	switch {
	case userRegister.Email == "":
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Email still empty",
			Data:    nil,
		})
	case userRegister.Password == "":
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password still empty",
			Data:    nil,
		})
	case userRegister.Fullname == "":
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Fullname still empty",
			Data:    nil,
		})
	case userRegister.Phone == "":
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Phone still empty",
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
