package routes

import (
	"assignment1/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.GET("/", controllers.HomeController)
	e.POST("/users/create", controllers.CreateUser)
	e.GET("/users", controllers.GetUsersPaginated)
	e.DELETE("/users/:id", controllers.DeleteUser)
}
