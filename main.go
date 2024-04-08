package main

import (
	"assignment1/configs"
	"assignment1/database"
	"assignment1/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	configs.Init()
	database.Init()
	routes.InitRoute(e)
	e.Logger.Fatal(e.Start(":1323"))
}
