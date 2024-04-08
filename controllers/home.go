package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeController(c echo.Context) error {
	return c.String(http.StatusOK, "Home Page")
}
