package cinemas

import (
	"booking-ticket/business/cinemas"
	"booking-ticket/controllers"
	"booking-ticket/controllers/cinemas/requests"
	"booking-ticket/controllers/cinemas/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CinemaController struct {
	CinemaUsecase cinemas.Usecase
}

func NewCinemaController(cinemaUseCase cinemas.Usecase) *CinemaController {
	return &CinemaController{
		CinemaUsecase: cinemaUseCase,
	}
}

func (cinemaController CinemaController) AddCinema(c echo.Context) error {
	cinemaAdd := requests.CinemaAdd{}
	c.Bind(&cinemaAdd)

	ctx := c.Request().Context()
	cinema, error := cinemaController.CinemaUsecase.AddCinema(ctx, cinemaAdd.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(cinema))
}

func (cinemaController CinemaController) ListCinema(c echo.Context) error {
	ctx := c.Request().Context()
	cinemas, error := cinemaController.CinemaUsecase.ListCinema(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, cinemas)
}

func (cinemaController CinemaController) CinemaDetail(c echo.Context) error {
	ctx := c.Request().Context()
	cinema, error := cinemaController.CinemaUsecase.CinemaDetail(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(cinema))
}
