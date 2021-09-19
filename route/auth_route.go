package route

import (
	"booking-ticket/controllers"

	"github.com/labstack/echo/v4"
)

func AuthRoute(auth *echo.Group) {
	auth.POST("auth/register", controllers.RegisterController)
	auth.POST("auth/login", controllers.LoginController)
}
