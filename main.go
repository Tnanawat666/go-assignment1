package main

// @title Assignment 1
// @version 1.0
// @description This is a sample server assignment 1 Internships.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:1323
// @BasePath /
// @schemes http

import (
	"assignment1/configs"
	"assignment1/database"
	"assignment1/routes"

	_ "assignment1/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	// * Echo instance
	e := echo.New()

	// * Configurations
	configs.Init()
	database.Init()
	routes.InitRoute(e)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// * Start server
	e.Logger.Fatal(e.Start(":1323"))
}
