package controllers

import (
	"booking-ticket/config"
	"booking-ticket/models/response"
	"booking-ticket/models/schedules"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateScheduleController(c echo.Context) error {
	var scheduleCreate schedules.ScheduleCreate
	c.Bind(&scheduleCreate)

	// validasi
	if scheduleCreate.MovieId == 0 || scheduleCreate.CinemaId == 0 || scheduleCreate.Price == 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Please input all field",
			Data:    nil,
		})
	}
	fmt.Println(scheduleCreate)

	var scheduleDB schedules.Schedule
	scheduleDB.MovieId = scheduleCreate.MovieId
	scheduleDB.CinemaId = scheduleCreate.CinemaId
	scheduleDB.Date = scheduleCreate.Date
	scheduleDB.Price = scheduleCreate.Price
	scheduleDB.Status = 1

	result := config.DB.Create(&scheduleDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error, Please try again later!",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Create Schedule Successfully",
		Data:    scheduleDB,
	})
}

func GetScheduleController(c echo.Context) error {
	var scheduleDB []schedules.Schedule
	var scheduleRes []schedules.ScheduleResponse

	// err := config.DB.Find(&scheduleDB).Error
	err := config.DB.Model(&scheduleDB).Select("movies.title ,schedules.date, schedules.price").Joins("left join movies on schedules.movie_id = movies.id").Scan(&scheduleRes).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Schedule List",
		Data:    scheduleRes,
	})
}
