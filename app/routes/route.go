package routes

import (
	"booking-ticket/controllers/admins"
	"booking-ticket/controllers/bookings"
	"booking-ticket/controllers/cinemas"
	"booking-ticket/controllers/movies"
	"booking-ticket/controllers/schedules"
	"booking-ticket/controllers/timeSchedules"
	"booking-ticket/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig              middleware.JWTConfig
	JwtConfigAdmin         middleware.JWTConfig
	UserController         users.UserController
	AdminController        admins.AdminController
	MovieController        movies.MovieController
	CinemaController       cinemas.CinemaController
	ScheduleController     schedules.ScheduleController
	TimeScheduleController timeSchedules.TimeScheduleController
	BookingController      bookings.BookingController
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
	e.GET("cinema", cl.CinemaController.ListCinema)
	e.GET("cinema/detail", cl.CinemaController.CinemaDetail)
	e.POST("cinema", cl.CinemaController.AddCinema)
}

func (cl *ControllerList) RouteSchedule(e *echo.Group) {
	e.GET("schedule", cl.ScheduleController.ListSchedule)
	e.GET("schedule/:id", cl.ScheduleController.DetailTimeSchedule)
	e.POST("schedule", cl.ScheduleController.AddSchedule)
}

func (cl *ControllerList) RouteTimeSchedule(e *echo.Group) {
	e.POST("schedule/time", cl.TimeScheduleController.AddScheduleTime)
}

func (cl *ControllerList) RouteBooking(e *echo.Group) {
	e.POST("booking", cl.BookingController.AddBooking)
}
