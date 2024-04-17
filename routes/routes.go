package routes

import (
	"assignment1/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.GET("/", controllers.HomeController)
	e.POST("/company/create/table", controllers.CreateCompanyTable)

	// Create
	e.POST("/user/create", controllers.CreateUser)
	e.POST("/users/create", controllers.CreateMultipleUsers)
	e.POST("/companies/create", controllers.CreateCompany)

	// Read
	e.GET("/users", controllers.GetUsersPaginated)
	e.GET("/users/products/order", controllers.GetUserProductOrder)

	// Update
	e.PUT("/user/:id", controllers.UpdateById)

	e.PUT("/users", controllers.UpdateMultiple)

	// Delete
	e.DELETE("/user/:id", controllers.DeleteUser)

	e.DELETE("/users/delete", controllers.DeleteMultipleUsers)
}
