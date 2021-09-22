package controllers

import (
	"booking-ticket/config"
	"booking-ticket/models/cinemas"
	"booking-ticket/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateCinemaController(c echo.Context) error {
	var cinemaCreate cinemas.CinemaCreate
	c.Bind(&cinemaCreate)

	// validasi
	if cinemaCreate.Name == "" || cinemaCreate.Location == "" || cinemaCreate.Maps == "" || cinemaCreate.Phone == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Please input all field",
			Data:    nil,
		})
	}

	var cinemaDB cinemas.Cinema
	cinemaDB.Name = cinemaCreate.Name
	cinemaDB.Location = cinemaCreate.Location
	cinemaDB.Maps = cinemaCreate.Maps
	cinemaDB.Phone = cinemaCreate.Phone
	cinemaDB.Status = 1

	result := config.DB.Create(&cinemaDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error, Please try again later!",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Create Cinema Successfully",
		Data:    cinemaDB,
	})
}

func GetCinemaController(c echo.Context) error {
	var cinemasDB []cinemas.Cinema

	err := config.DB.Find(&cinemasDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Cinema List",
		Data:    cinemasDB,
	})
}
