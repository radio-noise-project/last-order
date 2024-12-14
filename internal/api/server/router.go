package server

import (
	"github.com/labstack/echo"
	"github.com/radio-noise-project/last-order/internal/api/handler/container"
	"github.com/radio-noise-project/last-order/internal/api/handler/node"
	"github.com/radio-noise-project/last-order/internal/api/handler/runtime"
)

func router(e *echo.Echo) {
	// if you want to update api version, you should update here
	v0 := e.Group("/api/v0")

	runtime.Handler(v0)
	node.Handler(v0)
	container.Handler(v0)
}
