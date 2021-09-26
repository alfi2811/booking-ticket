package timeSchedules

import (
	"booking-ticket/business/timeSchedules"
	"booking-ticket/controllers"
	"booking-ticket/controllers/timeSchedules/requests"
	"booking-ticket/controllers/timeSchedules/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TimeScheduleController struct {
	TImeScheduleUsecase timeSchedules.Usecase
}

func NewTimeScheduleController(timeScheduleUsecase timeSchedules.Usecase) *TimeScheduleController {
	return &TimeScheduleController{
		TImeScheduleUsecase: timeScheduleUsecase,
	}
}

func (timeScheduleController TimeScheduleController) AddScheduleTime(c echo.Context) error {
	timeScheduleAdd := requests.TimeScheduleAdd{}
	c.Bind(&timeScheduleAdd)

	ctx := c.Request().Context()
	schedule, error := timeScheduleController.TImeScheduleUsecase.AddScheduleTime(ctx, timeScheduleAdd.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(schedule))
}
