package routes

import (
	"assignment1/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.GET("/", controllers.HomeController)

	// Create
	e.POST("/user/create", controllers.CreateUser)
	e.POST("/users/create", controllers.CreateMultipleUsers)

	// Read
	e.GET("/users", controllers.GetUsersPaginated)

	// Update
	e.PUT("/user/:id", controllers.UpdateById)

	e.PUT("/users", controllers.UpdateMultiple)

	// Delete
	e.DELETE("/user/:id", controllers.DeleteUser)

	e.POST("/users", controllers.DeleteMultipleUsers)
}
