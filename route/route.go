package route

import (
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	apiV1 := e.Group("api/v1/")
	AuthRoute(apiV1)
	UserRoute(apiV1)

	return e
}
