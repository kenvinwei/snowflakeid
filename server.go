package snowflakeid

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {

	e := echo.New()

	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	for _, v := range getRouteList() {
		e.GET(v.url, v.handler)
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
