package users

import (
	"booking-ticket/business/users"
	"booking-ticket/controllers"
	"booking-ticket/controllers/users/requests"
	"booking-ticket/controllers/users/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.Usecase
}

func NewUserController(userUseCase users.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (userController UserController) Login(c echo.Context) error {
	userLogin := requests.UserLogin{}
	c.Bind(&userLogin)

	ctx := c.Request().Context()
	user, error := userController.UserUseCase.Login(ctx, userLogin.Email, userLogin.Password)

	if error != nil {
		if error.Error() == "password wrong" {
			return controllers.NewErrorResponse(c, http.StatusForbidden, error)
		} else if error.Error() == "email empty" || error.Error() == "password empty" {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, error)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainLogin(user))
}

func (userController UserController) Register(c echo.Context) error {
	userRegister := requests.UserRegister{}
	c.Bind(&userRegister)

	ctx := c.Request().Context()
	user, error := userController.UserUseCase.Register(ctx, userRegister.ToDomain())

	if error != nil {
		if error.Error() == "password wrong" {
			return controllers.NewErrorResponse(c, http.StatusForbidden, error)
		} else if error.Error() == "please input all field" {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, error)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}

func (userController UserController) GetUser(c echo.Context) error {
	ctx := c.Request().Context()
	user, error := userController.UserUseCase.GetUser(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromListDomain(user))
}

func (userController UserController) ListUserBooking(c echo.Context) error {
	ctx := c.Request().Context()
	user, error := userController.UserUseCase.ListUserBooking(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, user)
}
