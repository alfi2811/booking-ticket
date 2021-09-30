package movies

import (
	"booking-ticket/business/movies"
	"booking-ticket/controllers"
	"booking-ticket/controllers/movies/requests"
	"booking-ticket/controllers/movies/responses"
	"net/http"
	"strconv"

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

	return controllers.NewSuccesResponse(c, responses.FromDomainDetail(movie))
}

func (movieController MovieController) ListMovie(c echo.Context) error {
	ctx := c.Request().Context()
	movies, error := movieController.MovieUsecase.ListMovie(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromListDomain(movies))
}

func (movieController MovieController) DetailMovie(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	movie, error := movieController.MovieUsecase.DetailMovie(ctx, id)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainDetail(movie))
}
