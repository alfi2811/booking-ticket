package admins

import (
	"booking-ticket/business/admins"
	"booking-ticket/controllers"
	"booking-ticket/controllers/admins/requests"
	"booking-ticket/controllers/admins/responses"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
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
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}

func (adminController AdminController) Register(c echo.Context) error {
	fmt.Println("Login")
	userRegister := requests.AdminRegister{}
	c.Bind(&userRegister)

	ctx := c.Request().Context()
	user, error := adminController.AdminUseCase.Register(ctx, userRegister.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}

func (adminController AdminController) GetAdmin(c echo.Context) error {
	ctx := c.Request().Context()
	user, error := adminController.AdminUseCase.GetAdmin(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, user)
}
