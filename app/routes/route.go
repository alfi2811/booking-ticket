package routes

import (
	"booking-ticket/controllers/admins"
	"booking-ticket/controllers/cinemas"
	"booking-ticket/controllers/movies"
	"booking-ticket/controllers/schedules"
	"booking-ticket/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig          middleware.JWTConfig
	JwtConfigAdmin     middleware.JWTConfig
	UserController     users.UserController
	AdminController    admins.AdminController
	MovieController    movies.MovieController
	CinemaController   cinemas.CinemaController
	ScheduleController schedules.ScheduleController
}

func (cl *ControllerList) RouteUsers(e *echo.Group) {
	e.POST("users/login", cl.UserController.Login)
	e.POST("users/register", cl.UserController.Register)
	e.GET("users", cl.UserController.GetUser, middleware.JWTWithConfig(cl.JwtConfig))
}

func (cl *ControllerList) RouteAdmins(e *echo.Group) {
	e.POST("admins/login", cl.AdminController.Login)
	e.POST("admins/register", cl.AdminController.Register)
	e.GET("admins", cl.AdminController.GetAdmin, middleware.JWTWithConfig(cl.JwtConfigAdmin))
}

func (cl *ControllerList) RouteMovies(e *echo.Group) {
	e.GET("movie", cl.MovieController.ListMovie)
	e.POST("movie", cl.MovieController.AddMovie)
}

func (cl *ControllerList) RouteCinemas(e *echo.Group) {
	e.GET("cinema", cl.CinemaController.ListMovie)
	e.POST("cinema", cl.CinemaController.AddMovie)
}

func (cl *ControllerList) RouteSchedule(e *echo.Group) {
	e.GET("schedule", cl.ScheduleController.ListSchedule)
	e.POST("schedule", cl.ScheduleController.AddSchedule)
}
