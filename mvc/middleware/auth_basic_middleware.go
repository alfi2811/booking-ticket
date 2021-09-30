package middleware

import (
	"booking-ticket/config"
	"booking-ticket/models/users"

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var user users.User

	err := config.DB.Where("email = ? AND password = ?", username, password).Find(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
