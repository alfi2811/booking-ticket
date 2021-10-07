package admins

import (
	"booking-ticket/business/admins"
	"booking-ticket/controllers"
	"booking-ticket/controllers/admins/requests"
	"booking-ticket/controllers/admins/responses"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AdminController struct {
	AdminUseCase admins.Usecase
}

func NewAdminController(adminUseCase admins.Usecase) *AdminController {
	return &AdminController{
		AdminUseCase: adminUseCase,
	}
}

func (adminController AdminController) Login(c echo.Context) error {
	adminLogin := requests.AdminLogin{}
	c.Bind(&adminLogin)

	ctx := c.Request().Context()
	user, error := adminController.AdminUseCase.Login(ctx, adminLogin.Email, adminLogin.Password)

	if error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) || error.Error() == "password wrong" {
			return controllers.NewErrorResponse(c, http.StatusForbidden, error)
		} else if error.Error() == "email empty" || error.Error() == "password empty" {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, error)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainLogin(user))
}

func (adminController AdminController) Register(c echo.Context) error {
	adminRegister := requests.AdminRegister{}
	c.Bind(&adminRegister)

	ctx := c.Request().Context()
	admin, error := adminController.AdminUseCase.Register(ctx, adminRegister.ToDomain())

	if error != nil {
		if error.Error() == "please input all field" {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, error)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(admin))
}

func (adminController AdminController) GetAdmin(c echo.Context) error {
	ctx := c.Request().Context()
	admin, error := adminController.AdminUseCase.GetAdmin(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromListDomain(admin))
}
