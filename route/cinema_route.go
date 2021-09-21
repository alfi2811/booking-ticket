package route

import (
	"booking-ticket/controllers"

	"github.com/labstack/echo/v4"
)

func CinemaRoute(auth *echo.Group) {
	auth.GET("cinema", controllers.GetCinemaController)
	auth.POST("cinema", controllers.CreateCinemaController)
}
