package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nekoze1210/echo-api-w3w/interfaces/api/server/router"
)

func RunServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.NewRouter(e)

	e.Logger.Fatal(e.Start(":1323"))
}
