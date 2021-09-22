package route

import (
	"booking-ticket/controllers"

	"github.com/labstack/echo/v4"
)

func ScheduleRoute(auth *echo.Group) {
	auth.GET("schedule", controllers.GetScheduleController)
	auth.POST("schedule", controllers.CreateScheduleController)
}
