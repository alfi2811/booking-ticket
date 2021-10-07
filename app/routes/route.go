package routes

import (
	middlewareApp "booking-ticket/app/middlewares"
	controller "booking-ticket/controllers"
	"booking-ticket/controllers/admins"
	"booking-ticket/controllers/bookings"
	"booking-ticket/controllers/cinemas"
	"booking-ticket/controllers/movies"
	"booking-ticket/controllers/schedules"
	"booking-ticket/controllers/timeSchedules"
	"booking-ticket/controllers/users"
	"errors"
	"fmt"
	"net/http"

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
	e.GET("users/booking", cl.UserController.ListUserBooking)
}

func (cl *ControllerList) RouteAdmins(e *echo.Group) {
	e.POST("admins/login", cl.AdminController.Login)
	e.POST("admins/register", cl.AdminController.Register)
	e.GET("admins", cl.AdminController.GetAdmin, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationAdmin())
}

func (cl *ControllerList) RouteMovies(e *echo.Group) {
	e.GET("movie", cl.MovieController.ListMovie, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("movie/:id", cl.MovieController.DetailMovie, middleware.JWTWithConfig(cl.JwtConfig))
	e.POST("movie", cl.MovieController.AddMovie, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationAdmin())
}

func (cl *ControllerList) RouteCinemas(e *echo.Group) {
	e.GET("cinema", cl.CinemaController.ListCinema, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("cinema/:id", cl.CinemaController.CinemaDetail, middleware.JWTWithConfig(cl.JwtConfig))
	e.POST("cinema", cl.CinemaController.AddCinema, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationAdmin())
}

func (cl *ControllerList) RouteSchedule(e *echo.Group) {
	e.GET("schedule", cl.ScheduleController.ListSchedule)
	e.GET("schedule/:id", cl.ScheduleController.DetailTimeSchedule)
	e.POST("schedule", cl.ScheduleController.AddSchedule, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationAdmin())
}

func (cl *ControllerList) RouteTimeSchedule(e *echo.Group) {
	e.POST("schedule/time", cl.TimeScheduleController.AddScheduleTime, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationAdmin())
}

func (cl *ControllerList) RouteBooking(e *echo.Group) {
	e.GET("booking", cl.BookingController.ListBooking, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("booking/:id", cl.BookingController.DetailBooking, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("booking/user/:id", cl.BookingController.ListBookingUser, middleware.JWTWithConfig(cl.JwtConfig))
	e.POST("booking", cl.BookingController.AddBooking, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationAdmin())
}

func RoleValidation(role string, userControler users.UserController) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			userRole := claims.Role

			if userRole == role {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, errors.New("forbidden roles"))
			}
		}
	}
}

func RoleValidationAdmin() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)
			fmt.Println("MASUK")
			fmt.Println(claims)
			userRole := claims.Role

			if userRole == "admin" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, errors.New("forbidden roles, Not Admin"))
			}
		}
	}
}
