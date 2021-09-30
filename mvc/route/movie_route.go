package route

import (
	"booking-ticket/controllers"

	"github.com/labstack/echo/v4"
)

func MovieRoute(auth *echo.Group) {
	auth.GET("movie", controllers.GetMovieController)
	auth.POST("movie", controllers.CreateMovieController)
}
