package controllers

import (
	"booking-ticket/config"
	"booking-ticket/models/movies"
	"booking-ticket/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateMovieController(c echo.Context) error {
	var movieCreate movies.MovieCreate
	c.Bind(&movieCreate)

	// validasi
	if movieCreate.Title == "" || movieCreate.Duration == "" || movieCreate.Players == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Please input all field",
			Data:    nil,
		})
	}

	var movieDB movies.Movie
	movieDB.Title = movieCreate.Title
	movieDB.Poster = movieCreate.Poster
	movieDB.ReleaseDate = movieCreate.ReleaseDate
	movieDB.Duration = movieCreate.Duration
	movieDB.MovieTrailer = movieCreate.MovieTrailer
	movieDB.Players = movieCreate.Players
	movieDB.Status = 1

	result := config.DB.Create(&movieDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error, Please try again later!",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Create Movie Successfully",
		Data:    movieDB,
	})
}

func GetMovieController(c echo.Context) error {
	var moviesDB []movies.Movie

	err := config.DB.Find(&moviesDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Movie List",
		Data:    moviesDB,
	})
}
