package runtime

import (
	"github.com/labstack/echo"
)

func Handler(api *echo.Group) {
	api.GET("/runtime/ping/:archtecture", getPing)
	api.GET("/runtime/version/:archtecture", getVersion)
}
