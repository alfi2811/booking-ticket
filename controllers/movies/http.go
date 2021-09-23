package movies

import (
	"booking-ticket/business/movies"
	"booking-ticket/controllers"
	"booking-ticket/controllers/movies/requests"
	"booking-ticket/controllers/movies/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MovieController struct {
	MovieUsecase movies.Usecase
}

func NewMovieController(movieUseCase movies.Usecase) *MovieController {
	return &MovieController{
		MovieUsecase: movieUseCase,
	}
}

func (movieController MovieController) AddMovie(c echo.Context) error {
	movieAdd := requests.MovieAdd{}
	c.Bind(&movieAdd)

	ctx := c.Request().Context()
	movie, error := movieController.MovieUsecase.AddMovie(ctx, movieAdd.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(movie))
}
