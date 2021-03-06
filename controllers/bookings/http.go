package bookings

import (
	"booking-ticket/business/bookings"
	"booking-ticket/controllers"
	"booking-ticket/controllers/bookings/requests"
	"booking-ticket/controllers/bookings/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookingController struct {
	BookingUsecase bookings.Usecase
}

func NewBookingController(bookingUseCase bookings.Usecase) *BookingController {
	return &BookingController{
		BookingUsecase: bookingUseCase,
	}
}

func (bookingController BookingController) AddBooking(c echo.Context) error {
	bookingAdd := requests.BookingAdd{}
	c.Bind(&bookingAdd)

	ctx := c.Request().Context()
	booking, error := bookingController.BookingUsecase.AddBooking(ctx, bookingAdd.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromAddDomain(booking))
}

func (bookingController BookingController) ListBooking(c echo.Context) error {
	ctx := c.Request().Context()
	bookings, error := bookingController.BookingUsecase.ListBooking(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromListDomain(bookings))
}

func (bookingController BookingController) ListBookingUser(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	bookings, error := bookingController.BookingUsecase.ListBookingUser(ctx, id)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromListDomainBookingUser(bookings))
}

func (bookingController BookingController) DetailBooking(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	bookings, error := bookingController.BookingUsecase.DetailBooking(ctx, id)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainDetail(bookings))
}
