package routes

import (
	"booking-ticket/controllers/cinemas"
	"booking-ticket/controllers/movies"
	"booking-ticket/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController   users.UserController
	MovieController  movies.MovieController
	CinemaController cinemas.CinemaController
}

func (cl *ControllerList) RouteUsers(e *echo.Echo) {
	e.POST("users/login", cl.UserController.Login)
	e.POST("users/register", cl.UserController.Register)
	e.GET("users", cl.UserController.GetUser)
}

func (cl *ControllerList) RouteMovies(e *echo.Echo) {
	e.POST("movie", cl.MovieController.AddMovie)
}

func (cl *ControllerList) RouteCinemas(e *echo.Echo) {
	e.POST("cinema", cl.CinemaController.AddMovie)
}
