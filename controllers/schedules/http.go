package schedules

import (
	"booking-ticket/business/schedules"
	"booking-ticket/controllers"
	"booking-ticket/controllers/schedules/requests"
	"booking-ticket/controllers/schedules/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ScheduleController struct {
	ScheduleUsecase schedules.Usecase
}

func NewScheduleController(scheduleUsecase schedules.Usecase) *ScheduleController {
	return &ScheduleController{
		ScheduleUsecase: scheduleUsecase,
	}
}

func (scheduleController ScheduleController) AddSchedule(c echo.Context) error {
	scheduleAdd := requests.ScheduleAdd{}
	c.Bind(&scheduleAdd)

	ctx := c.Request().Context()
	schedule, error := scheduleController.ScheduleUsecase.AddSchedule(ctx, scheduleAdd.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(schedule))
}

func (scheduleController ScheduleController) ListSchedule(c echo.Context) error {
	ctx := c.Request().Context()
	schedules, error := scheduleController.ScheduleUsecase.ListSchedule(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromListDomainJoin(schedules))
}

func (scheduleController ScheduleController) DetailTimeSchedule(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	schedules, error := scheduleController.ScheduleUsecase.DetailTimeSchedule(ctx, id)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDetailJoin(schedules))
}
