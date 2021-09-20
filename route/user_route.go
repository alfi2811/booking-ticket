package route

import (
	"booking-ticket/constants"
	"booking-ticket/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserRoute(user *echo.Group) {
	// user.Use(middleware.BasicAuth(m.BasicAuthDB))
	user.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	user.GET("users", controllers.GetUserController)
	user.GET("users/:id", controllers.GetUserDetailController)
	user.PUT("users/:id", controllers.UpdateUserController)
	user.DELETE("users/:id", controllers.DeleteUserController)
}
