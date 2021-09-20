package route

import (
	"booking-ticket/controllers"

	"github.com/labstack/echo/v4"
)

func AuthRoute(auth *echo.Group) {
	auth.POST("auth/register/user", controllers.RegisterController)
	auth.POST("auth/login/user", controllers.LoginController)
	auth.POST("auth/register/admin", controllers.RegisterAdminController)
	auth.POST("auth/login/admin", controllers.LoginAdminController)
}
