package route

import (
	m "booking-ticket/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	apiV1 := e.Group("api/v1/")
	AuthRoute(apiV1)
	m.LogMiddleware(e)
	UserRoute(apiV1)
	MovieRoute(apiV1)
	CinemaRoute(apiV1)

	return e
}
