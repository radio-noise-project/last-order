package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/radio-noise-project/last-order/internal/api/handler"
)

func Router() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define APIs
	e.GET("/v0/runtime/version", handler.OutputLastOrderVersion)

	e.Logger.Fatal(e.Start(":8080"))
}
