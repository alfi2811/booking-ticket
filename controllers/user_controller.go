package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"booking-ticket/config"
	"booking-ticket/models/response"
	"booking-ticket/models/users"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	var users []users.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    users,
	})
}

func GetUserDetailController(c echo.Context) error {
	var users users.User

	id, _ := strconv.Atoi(c.Param("id"))

	result := config.DB.First(&users, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return c.JSON(http.StatusBadRequest, response.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: result.Error.Error(),
				Data:    nil,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: result.Error.Error(),
				Data:    nil,
			})
		}
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Detail Data User",
		Data:    users,
	})
}

func UpdateUserController(c echo.Context) error {
	var userUpdate users.UserUpdate
	c.Bind(&userUpdate)

	id, _ := strconv.Atoi(c.Param("id"))

	var userDB users.User

	fmt.Println(userUpdate)

	result := config.DB.Model(&userDB).Where("id = ?", id).Updates(
		map[string]interface{}{
			"id":       id,
			"email":    userUpdate.Email,
			"password": userUpdate.Password,
			"fullname": userUpdate.Fullname,
			"gender":   userUpdate.Gender,
			"phone":    userUpdate.Phone,
		})
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Error, Please try again later!",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Update Data User Successfully",
		Data:    userDB,
	})
}

func DeleteUserController(c echo.Context) error {
	var users users.User

	id, _ := strconv.Atoi(c.Param("id"))

	result := config.DB.Where("id = ?", id).Delete(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: result.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Delete Data User Successfull",
		Data:    users,
	})
}
