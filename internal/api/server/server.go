package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	// Create new echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	router(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
