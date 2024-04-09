package main

import (
	"assignment1/configs"
	"assignment1/database"
	"assignment1/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// * Echo instance
	e := echo.New()

	// * Configurations
	configs.Init()
	database.Init()
	routes.InitRoute(e)

	// * Start server
	e.Logger.Fatal(e.Start(":1323"))
}
