package route

import (
	"booking-ticket/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(user *echo.Group) {
	user.GET("users", controllers.GetUserController)
	user.GET("users/:id", controllers.GetUserDetailController)
	user.PUT("users/:id", controllers.UpdateUserController)
	user.DELETE("users/:id", controllers.DeleteUserController)
}
