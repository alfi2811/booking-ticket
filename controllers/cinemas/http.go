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

func (cinemaController CinemaController) AddMovie(c echo.Context) error {
	cinemaAdd := requests.CinemaAdd{}
	c.Bind(&cinemaAdd)

	ctx := c.Request().Context()
	cinema, error := cinemaController.CinemaUsecase.AddCinema(ctx, cinemaAdd.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(cinema))
}
